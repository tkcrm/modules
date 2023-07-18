package taskmanager

import (
	"context"
	"fmt"
)

type HandlerFunc func(ctx context.Context, t *Task) error

type WorkerHandlers map[string]HandlerFunc

func (s *service) RegisterWorkerHandlers(handlers WorkerHandlers) error {
	for name, handleFn := range handlers {
		if name == "" {
			return fmt.Errorf("register worker handlers error: empty name")
		}

		if handleFn == nil {
			return fmt.Errorf("register \"%s\" worker handlers error: empty handle function", name)
		}

		if err := s.RegisterWorkerHandler(name, handleFn); err != nil {
			return fmt.Errorf("register \"%s\" worker handler error: %w", name, err)
		}
	}

	return nil
}
