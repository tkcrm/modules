package taskmanager

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/goccy/go-json"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tkcrm/modules/pkg/logger"
	"github.com/vgarvardt/gue/v5"
	"github.com/vgarvardt/gue/v5/adapter/pgxv5"
	"golang.org/x/sync/errgroup"
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

	ctx         context.Context
	stopChan    chan struct{}
	stoppedChan chan struct{}
	started     bool
	blocked     bool

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
func (s *service) Name() string { return "taskmanager" }

// Start task manager
func (s *service) Start(ctx context.Context) error {
	// skip if service already started
	if s.IsStarted() {
		return nil
	}

	s.setIsStarted(true)
	defer s.setIsStarted(false)

	s.ctx = ctx

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

	if _, ok := s.workerHandlers[name]; ok {
		return fmt.Errorf("worker handler with name [%s] already exists", name)
	}

	s.mu.Lock()
	s.workerHandlers[name] = func(ctx context.Context, j *gue.Job) error {
		return fn(ctx, &Task{j})
	}
	s.mu.Unlock()

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

	var options options
	for _, opt := range opts {
		opt(&options)
	}

	if taskType == "" {
		return fmt.Errorf("empty task type")
	}

	if options.queueName == "" {
		options.queueName = DefaultQueueName
	}

	var payloadBytes []byte
	if payload != nil {
		switch t := payload.(type) {
		case []byte:
			payloadBytes = t
		case string:
			payloadBytes = []byte(t)
		default:
			if options.jsonPayload {
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
		Queue: options.queueName,
	}

	if !options.runAt.IsZero() {
		t.RunAt = options.runAt
	}

	if options.priority != 0 {
		t.Priority = gue.JobPriority(options.priority)
	}

	if err := s.gc.Enqueue(s.ctx, t); err != nil {
		if strings.Contains(err.Error(), "SQLSTATE 23505") {
			return nil
		}
		return fmt.Errorf("failed to enqueue job: %w", err)
	}

	return nil
}

func (s *service) CleanTableIndexes(ctx context.Context) error {
	s.mu.Lock()
	if s.blocked {
		s.mu.Unlock()
		return nil
	}
	s.blocked = true
	s.mu.Unlock()

	defer func() {
		s.mu.Lock()
		s.blocked = false
		s.mu.Unlock()
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
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.started
}

func (s *service) setIsStarted(v bool) {
	s.mu.Lock()
	s.started = v
	s.mu.Unlock()
}
