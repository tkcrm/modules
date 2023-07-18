package limiter

import (
	"fmt"
	"sync"

	"github.com/tkcrm/modules/pkg/logger"
)

type ILimiter interface {
	RegisterServices(services ...*Service) error
	GetService(name string) (*Service, error)
	GetServices() []*Service
}

type limiter struct {
	logger logger.Logger
	config Config

	redisClient IRedisClient

	mu       sync.Mutex
	services []*Service
}

func New(logger logger.Logger, config Config, redisClient IRedisClient, opts ...LimiterOption) (ILimiter, error) {
	if redisClient == nil {
		return nil, fmt.Errorf("redis client is nil")
	}

	l := &limiter{
		logger:      logger,
		config:      config,
		redisClient: redisClient,
		services:    []*Service{},
	}

	options := limiterOptions{}
	for _, opt := range opts {
		opt(&options)
	}

	return l, nil
}

// RegisterService register new service
func (l *limiter) RegisterServices(services ...*Service) error {
	l.mu.Lock()

	for _, svc := range services {
		// validate all new services
		if err := svc.Validate(); err != nil {
			return fmt.Errorf("validate service error: %w", err)
		}

		// check for unique service
		for i := range l.services {
			if l.services[i].Name == svc.Name {
				return fmt.Errorf("service with name %s already registered", svc.Name)
			}
		}

		// init new service
		if err := svc.initService(l.redisClient, l.config.CachePrefix); err != nil {
			return fmt.Errorf("init srvice %s error: %w", svc.Name, err)
		}

		// add new service to slice
		l.services = append(l.services, svc)
	}

	l.mu.Unlock()

	return nil
}

func (l *limiter) GetService(name string) (*Service, error) {
	for i := range l.services {
		if l.services[i].Name == name {
			return l.services[i], nil
		}
	}
	return nil, fmt.Errorf("failed to get service by name %s", name)
}

func (l *limiter) GetServices() []*Service {
	return l.services
}
