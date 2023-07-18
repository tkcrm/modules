package limiter

import (
	"context"
	"fmt"
	"time"

	externalLimiter "github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/redis"
)

const defaultCacheKey = "req"

type Service struct {
	// service name. should be unique
	Name string
	// period
	period time.Duration
	// limit per period
	limit uint64
	// formated limit: 10-S / 100-M / 1000-H / 10000-D
	formattedLimit string

	globalCachePrefix string

	externalLimiter *externalLimiter.Limiter
}

// NewService return new service instance
func NewService(name string, opts ...ServiceInitOption) *Service {
	s := &Service{
		Name: name,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

// Validate service
func (s *Service) Validate() error {
	if s.Name == "" {
		return fmt.Errorf("empty service name")
	}

	if s.formattedLimit == "" {
		if s.period == 0 {
			return fmt.Errorf("empty period")
		}

		if s.limit == 0 {
			return fmt.Errorf("empty limit")
		}

		if s.period == 0 && s.limit == 0 {
			return fmt.Errorf("empty formatted limit and period limit")
		}
	}

	return nil
}

func (s *Service) initService(redisClient IRedisClient, cachePrefix string) error {
	store, err := redis.NewStoreWithOptions(redisClient, externalLimiter.StoreOptions{
		Prefix: s.getCacheKey(),
	})
	if err != nil {
		return fmt.Errorf("init redis store error: %w", err)
	}

	rate := externalLimiter.Rate{
		Period: s.period,
		Limit:  int64(s.limit),
	}

	if s.formattedLimit != "" {
		parsedData, err := externalLimiter.NewRateFromFormatted(s.formattedLimit)
		if err != nil {
			return err
		}

		rate = parsedData
		s.period = parsedData.Period
		s.limit = uint64(parsedData.Limit)
	}

	s.externalLimiter = externalLimiter.New(store, rate)
	s.globalCachePrefix = cachePrefix

	return nil
}

type serviceLimitStats struct {
	Limit     int64
	Remaining int64
	Reset     int64
	Reached   bool
}

func (s *Service) getCacheKey() string {
	if s.globalCachePrefix != "" {
		return fmt.Sprintf("%s-limiter-%s", s.globalCachePrefix, s.Name)
	}
	return fmt.Sprintf("limiter-%s", s.Name)
}

func parseServiceOpts(opts ...ServiceMethodOption) serviceOptions {
	options := serviceOptions{}
	for _, opt := range opts {
		opt(&options)
	}

	if options.cacheKey == "" {
		options.cacheKey = defaultCacheKey
	}

	return options
}

// Increment custom integer value
func (s *Service) Increment(ctx context.Context, count uint64, opts ...ServiceMethodOption) (*serviceLimitStats, error) {
	options := parseServiceOpts(opts...)

	stats, err := s.externalLimiter.Increment(ctx, options.cacheKey, int64(count))
	if err != nil {
		return nil, err
	}

	res := serviceLimitStats(stats)

	return &res, nil
}

// Get increment one request and return new stats
func (s *Service) Get(ctx context.Context, opts ...ServiceMethodOption) (*serviceLimitStats, error) {
	options := parseServiceOpts(opts...)

	stats, err := s.externalLimiter.Get(ctx, options.cacheKey)
	if err != nil {
		return nil, err
	}

	res := serviceLimitStats(stats)

	return &res, nil
}

// Peek return current limit stats
func (s *Service) Peek(ctx context.Context, opts ...ServiceMethodOption) (*serviceLimitStats, error) {
	options := parseServiceOpts(opts...)

	stats, err := s.externalLimiter.Peek(ctx, options.cacheKey)
	if err != nil {
		return nil, err
	}

	res := serviceLimitStats(stats)

	return &res, nil
}

// Reset data
func (s *Service) Reset(ctx context.Context, opts ...ServiceMethodOption) (*serviceLimitStats, error) {
	options := parseServiceOpts(opts...)

	stats, err := s.externalLimiter.Reset(ctx, options.cacheKey)
	if err != nil {
		return nil, err
	}

	res := serviceLimitStats(stats)

	return &res, nil
}

// IsReached return is limit reached
func (s *Service) IsReached(ctx context.Context, opts ...ServiceMethodOption) (bool, error) {
	options := parseServiceOpts(opts...)

	stats, err := s.externalLimiter.Peek(ctx, options.cacheKey)
	if err != nil {
		return false, err
	}

	return stats.Reached, nil
}
