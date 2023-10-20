package taskmanager

import (
	"context"

	"github.com/hibiken/asynq"
)

type ITaskmanager interface {
	Name() string
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	RegisterWorkerHandler(name string, fn HandlerFunc) error
	RegisterWorkerHandlers(handlers WorkerHandlers) error
	AddTask(taskType string, payload any, opts ...Option) error
}

type Task struct {
	*asynq.Task
}

type HandlerFunc func(ctx context.Context, t *Task) error

type WorkerHandlers map[string]HandlerFunc
