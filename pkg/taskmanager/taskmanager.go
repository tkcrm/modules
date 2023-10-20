package taskmanager

import (
	"context"
	"errors"
	"fmt"

	"github.com/goccy/go-json"
	"github.com/hibiken/asynq"
)

type taskmanager struct {
	logger logger
	config Config

	srv      *asynq.Server
	client   *asynq.Client
	handlers *asynq.ServeMux
}

func New(logger logger, config Config) ITaskmanager {
	if config.ServerConfig.ShutdownTimeout == 0 {
		config.ServerConfig.ShutdownTimeout = defaultGracefulShutdownTimeDuration
	}
	config.ServerConfig.Logger = logger

	redisConfig := asynq.RedisClientOpt{
		Addr:     config.RedisConfig.Addr,
		Username: config.RedisConfig.Username,
		Password: config.RedisConfig.Password,
		DB:       config.RedisConfig.DB,
	}

	tm := &taskmanager{
		logger: logger,
		config: config,
		srv: asynq.NewServer(
			redisConfig,
			config.ServerConfig.Config,
		),
		client:   asynq.NewClient(redisConfig),
		handlers: asynq.NewServeMux(),
	}

	return tm
}

// Name return name of the service
func (s taskmanager) Name() string {
	if s.config.Name == "" {
		return defaultName
	}
	return s.config.Name
}

// Start task manager server
func (s *taskmanager) Start(ctx context.Context) error {
	if err := s.srv.Start(s.handlers); err != nil {
		return fmt.Errorf("failed to run task manager server: %w", err)
	}
	return nil
}

// Stop task manager server and client
func (s *taskmanager) Stop(ctx context.Context) error {
	// stop client
	if err := s.client.Close(); err != nil {
		s.logger.Errorf("failed to stop [%s] client: %s", s.Name(), err)
	}

	// stop server
	s.srv.Stop()

	// graceful shutdown the server
	s.srv.Shutdown()

	return nil
}

// RegisterWorkerHandler register new handler for task
func (s *taskmanager) RegisterWorkerHandler(name string, fn HandlerFunc) error {
	if name == "" {
		return fmt.Errorf("empty handler name")
	}

	if fn == nil {
		return fmt.Errorf("empty handler func")
	}

	s.handlers.HandleFunc(name, func(ctx context.Context, t *asynq.Task) error {
		return fn(ctx, &Task{t})
	})

	return nil
}

// RegisterWorkerHandlers register new handlers for tasks
func (s *taskmanager) RegisterWorkerHandlers(handlers WorkerHandlers) error {
	for name, handleFn := range handlers {
		if name == "" {
			return fmt.Errorf("register worker handlers error: empty name")
		}

		if handleFn == nil {
			return fmt.Errorf("register [%s] worker handlers error: empty handle function", name)
		}

		if err := s.RegisterWorkerHandler(name, handleFn); err != nil {
			return fmt.Errorf("register [%s] worker handler error: %w", name, err)
		}
	}

	return nil
}

// AddTask add new task
func (s *taskmanager) AddTask(name string, payload any, opts ...Option) error {
	o := &options{
		queue: defaultQueueName,
	}

	for _, opt := range opts {
		opt(o)
	}

	optsAsync := []asynq.Option{}
	if o.retry != 0 {
		optsAsync = append(optsAsync, asynq.MaxRetry(o.retry))
	}

	if o.queue != "" {
		optsAsync = append(optsAsync, asynq.Queue(o.queue))
	}

	if o.taskID != "" && !s.config.UniqueTasks {
		optsAsync = append(optsAsync, asynq.TaskID(o.taskID))
	}

	if o.timeout != 0 {
		optsAsync = append(optsAsync, asynq.Timeout(o.timeout))
	}

	if !o.deadline.IsZero() {
		optsAsync = append(optsAsync, asynq.Deadline(o.deadline))
	}

	if o.uniqueTTL != 0 {
		optsAsync = append(optsAsync, asynq.Unique(o.uniqueTTL))
	}

	if !o.processAt.IsZero() {
		optsAsync = append(optsAsync, asynq.ProcessAt(o.processAt))
	}

	if o.retention != 0 {
		optsAsync = append(optsAsync, asynq.Retention(o.retention))
	}

	if o.group != "" {
		optsAsync = append(optsAsync, asynq.Group(o.group))
	}

	var payloadBytes []byte
	if payload != nil {
		switch t := payload.(type) {
		case []byte:
			payloadBytes = t
		case string:
			payloadBytes = []byte(t)
		default:
			if o.jsonPayload {
				pb, err := json.Marshal(payload)
				if err != nil {
					return fmt.Errorf("marshal json payload error: %w", err)
				}
				payloadBytes = pb
			} else {
				return fmt.Errorf("bad payload type \"%T\" for task %s", t, name)
			}
		}
	}

	if s.config.UniqueTasks {
		taskID := genTaskID(fmt.Sprintf("%s-%s", o.queue, name), payloadBytes)
		optsAsync = append(optsAsync, asynq.TaskID(taskID))
	}

	task := asynq.NewTask(name, payloadBytes, optsAsync...)
	if _, err := s.client.Enqueue(task); err != nil {
		if errors.Is(err, asynq.ErrDuplicateTask) ||
			errors.Is(err, asynq.ErrTaskIDConflict) {
			return nil
		}
		return fmt.Errorf("enqueue new task error: %w", err)
	}

	return nil
}
