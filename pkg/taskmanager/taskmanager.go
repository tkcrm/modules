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
)

const DefaultQueueName = "default"

type ITaskmanager interface {
	Name() string
	Stop(ctx context.Context) error
	RegisterWorkerHandler(name string, fn HandlerFunc) error
	RegisterWorkerHandlers(handlers WorkerHandlers) error
	AddTask(taskType string, payload any, opts ...Option) error
	StartWorkers(poolSize int, queueName string)
}

type service struct {
	gc     *gue.Client
	logger logger.Logger
	db     *pgxpool.Pool

	stopFn context.CancelFunc
	ctx    context.Context

	mu      sync.RWMutex
	workMap gue.WorkMap
}

func New(ctx context.Context, l logger.Logger, psqlPool *pgxpool.Pool) (ITaskmanager, error) {
	poolAdapter := pgxv5.NewConnPool(psqlPool)
	gc, err := gue.NewClient(poolAdapter)
	if err != nil {
		return nil, fmt.Errorf("failed to create gue client: %w", err)
	}

	ctx, shutdown := context.WithCancel(ctx)

	return &service{
		logger:  l,
		gc:      gc,
		db:      psqlPool,
		workMap: make(gue.WorkMap),
		stopFn:  shutdown,
		ctx:     ctx,
	}, nil
}

func (s *service) RegisterWorkerHandler(name string, fn HandlerFunc) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("register \"%s\" worker handler error: %w", name, err)
		}
	}()

	if name == "" {
		return fmt.Errorf("empty handler name")
	}

	if fn == nil {
		return fmt.Errorf("empty handler func")
	}

	s.mu.Lock()
	s.workMap[name] = func(ctx context.Context, j *gue.Job) error {
		return fn(ctx, &Task{j})
	}
	s.mu.Unlock()

	return nil
}

func (s *service) initWorkers(poolSize int, queueName string) error {
	if poolSize == 0 {
		poolSize = 1
	}

	if queueName == "" {
		queueName = DefaultQueueName
	}

	workers, err := gue.NewWorkerPool(s.gc, s.workMap, poolSize,
		gue.WithPoolQueue(queueName),
		gue.WithPoolPollInterval(time.Second),
	)
	if err != nil {
		return fmt.Errorf("NewWorkerPool error: %w", err)
	}

	go func() {
		if err := workers.Run(s.ctx); err != nil {
			s.logger.Errorf("worker run error: %v", err)
		}
	}()

	return nil
}

func (s *service) StartWorkers(poolSize int, queueName string) {
	if err := s.initWorkers(poolSize, queueName); err != nil {
		s.logger.Errorf("initWorkers error: %v", err)
	}

	s.logger.Infof("task manager was started for queue: %s", queueName)

	<-s.ctx.Done()
}

func (s *service) AddTask(taskType string, payload any, opts ...Option) (err error) {
	defer func() {
		if err != nil && taskType != "" {
			err = fmt.Errorf("add task \"%s\" error: %w", taskType, err)
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

func (s *service) Name() string { return "taskmanager" }

func (s *service) Stop(ctx context.Context) error {
	if s.stopFn != nil {
		s.stopFn()
		s.ctx = nil
		s.stopFn = nil
	}
	return nil
}
