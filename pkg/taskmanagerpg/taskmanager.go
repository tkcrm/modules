package taskmanagerpg

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/goccy/go-json"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tkcrm/modules/pkg/logger"
	"github.com/vgarvardt/gue/v5"
	"github.com/vgarvardt/gue/v5/adapter/pgxv5"
)

const (
	DefaultQueueName                    = "default"
	defaultGracefulShutdownTimeDuration = time.Second * 20
)

type ITaskmanager interface {
	Name() string
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	RegisterWorkerHandler(name string, fn HandlerFunc) error
	RegisterWorkerHandlers(handlers WorkerHandlers) error
	RegisterWorker(poolSize int, queueName string)
	AddTask(taskType string, payload any, opts ...Option) error
	CleanTableIndexes(ctx context.Context) error
	IsStarted() bool
}

type service struct {
	logger logger.Logger
	gc     *gue.Client
	db     *pgxpool.Pool

	muCtx sync.Mutex
	ctx   context.Context

	stopChan    chan struct{}
	stoppedChan chan struct{}

	muStarted sync.Mutex
	started   bool

	muBlocked sync.Mutex
	blocked   bool

	mu             sync.RWMutex
	workerHandlers gue.WorkMap
	workers        []worker
}

func New(l logger.Logger, psqlPool *pgxpool.Pool) (ITaskmanager, error) {
	poolAdapter := pgxv5.NewConnPool(psqlPool)
	gc, err := gue.NewClient(poolAdapter)
	if err != nil {
		return nil, fmt.Errorf("failed to create gue client: %w", err)
	}

	return &service{
		logger:         l,
		gc:             gc,
		db:             psqlPool,
		workerHandlers: make(gue.WorkMap),
		workers:        make([]worker, 0),
	}, nil
}

// Name return name of service
func (s *service) Name() string { return "taskmanagerpg" }

// Start task manager
func (s *service) Start(ctx context.Context) error {
	// skip if service already started
	if s.IsStarted() {
		return nil
	}

	s.setIsStarted(true)
	defer s.setIsStarted(false)

	s.muCtx.Lock()
	s.ctx = ctx
	s.muCtx.Unlock()

	// local context for graceful shutdown
	gracafulCtx, shutdown := context.WithCancel(context.Background())
	defer shutdown()

	g, groupCtx := errgroup.WithContext(gracafulCtx)
	for _, w := range s.workers {
		gueWorker, err := gue.NewWorkerPool(
			s.gc, s.workerHandlers, w.poolSize,
			gue.WithPoolQueue(w.queueName),
			gue.WithPoolPollInterval(time.Second),
		)
		if err != nil {
			return fmt.Errorf("gue NewWorkerPool error: %w", err)
		}

		g.Go(func() error {
			if err := gueWorker.Run(groupCtx); err != nil {
				return err
			}

			return nil
		})
	}

	s.logger.Info("task manager started")

	// waiting for stop signal
	s.stopChan = make(chan struct{}, 1)

	// signaling that service was stopped
	s.stoppedChan = make(chan struct{}, 1)
	defer func() {
		s.stoppedChan <- struct{}{}
	}()

	// errors workers channel
	groupErrChan := make(chan error, 1)

	// waiting for errors from workers
	go func() {
		groupErrChan <- g.Wait()
	}()

	select {
	case err := <-groupErrChan:
		if err != nil {
			return err
		}
	case <-s.stopChan:
	case <-gracafulCtx.Done():
	case <-ctx.Done():
	}

	shutdown()

	select {
	case <-time.After(defaultGracefulShutdownTimeDuration):
	case <-gracafulCtx.Done():
	}

	s.logger.Info("task manager stopped")

	return nil
}

// Stop task manager
func (s *service) Stop(ctx context.Context) error {
	// skip if service not running
	if !s.IsStarted() {
		return nil
	}

	// send stop signal
	s.stopChan <- struct{}{}

	// waiting for grace stop all tasks
	select {
	case <-ctx.Done():
	case <-s.stoppedChan:
	}

	return nil
}

// RegisterWorkerHandler register a new worker handler
func (s *service) RegisterWorkerHandler(name string, fn HandlerFunc) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("register [%s] worker handler error: %w", name, err)
		}
	}()

	if name == "" {
		return fmt.Errorf("empty handler name")
	}

	if fn == nil {
		return fmt.Errorf("empty handler func")
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.workerHandlers[name]; ok {
		return fmt.Errorf("worker handler with name [%s] already exists", name)
	}

	s.workerHandlers[name] = func(ctx context.Context, j *gue.Job) error {
		return fn(ctx, &Task{j})
	}

	return nil
}

// RegisterWorker register a new worker
func (s *service) RegisterWorker(poolSize int, queueName string) {
	if poolSize == 0 {
		poolSize = 1
	}

	if queueName == "" {
		queueName = DefaultQueueName
	}

	w := worker{
		poolSize:  poolSize,
		queueName: queueName,
	}

	s.mu.Lock()
	s.workers = append(s.workers, w)
	s.mu.Unlock()

	s.logger.Debugf("task manager registered worker pool size %d for queue: %s", poolSize, queueName)
}

// AddTask add a new task to queue
func (s *service) AddTask(taskType string, payload any, opts ...Option) (err error) {
	if !s.IsStarted() {
		return nil
	}

	defer func() {
		if err != nil && taskType != "" {
			err = fmt.Errorf("add task [%s] error: %w", taskType, err)
		}
	}()

	var options Options
	for _, opt := range opts {
		opt(&options)
	}

	if taskType == "" {
		return fmt.Errorf("empty task type")
	}

	if options.QueueName == "" {
		options.QueueName = DefaultQueueName
	}

	var payloadBytes []byte
	if payload != nil {
		switch t := payload.(type) {
		case []byte:
			payloadBytes = t
		case string:
			payloadBytes = []byte(t)
		default:
			if options.JSONPayload {
				pb, err := json.Marshal(payload)
				if err != nil {
					return err
				}
				payloadBytes = pb
			} else {
				return fmt.Errorf("bad payload type \"%T\" for task %s", t, taskType)
			}
		}
	}

	t := &gue.Job{
		Type:  taskType,
		Args:  payloadBytes,
		Queue: options.QueueName,
	}

	if !options.RunAt.IsZero() {
		t.RunAt = options.RunAt
	}

	if options.Priority != 0 {
		t.Priority = gue.JobPriority(options.Priority)
	}

	s.muCtx.Lock()
	ctx := s.ctx
	s.muCtx.Unlock()

	if err := s.gc.Enqueue(ctx, t); err != nil {
		if strings.Contains(err.Error(), "SQLSTATE 23505") {
			return nil
		}
		return fmt.Errorf("failed to enqueue job: %w", err)
	}

	return nil
}

func (s *service) CleanTableIndexes(ctx context.Context) error {
	s.muBlocked.Lock()
	if s.blocked {
		s.muBlocked.Unlock()
		return nil
	}
	s.blocked = true
	s.muBlocked.Unlock()

	defer func() {
		s.muBlocked.Lock()
		s.blocked = false
		s.muBlocked.Unlock()
	}()

	if !s.IsStarted() {
		return nil
	}

	s.logger.Info("start cleaning the job table")

	// stop task manager
	if err := s.Stop(ctx); err != nil {
		return fmt.Errorf("failed to stop task manager: %w", err)
	}

	now := time.Now()

	// clean table
	if _, err := s.db.Exec(ctx, "VACUUM FULL VERBOSE gue_jobs;"); err != nil {
		return fmt.Errorf("db execute error: %w", err)
	}

	s.logger.Infof("table gue_job successfully cleaned in: %s", time.Since(now))

	// start task manager
	go func() {
		if err := s.Start(ctx); err != nil {
			s.logger.Fatalf("failed to start task manager: %s", err)
		}
	}()

	return nil
}

func (s *service) IsStarted() bool {
	s.muStarted.Lock()
	defer s.muStarted.Unlock()
	return s.started
}

func (s *service) setIsStarted(v bool) {
	s.muStarted.Lock()
	s.started = v
	s.muStarted.Unlock()
}
