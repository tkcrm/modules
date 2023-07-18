package limiter

import "time"

/* ServiceInitOption */
type ServiceInitOption func(*Service)

// You can also use the simplified format "<limit>-<period>"", with the given
// periods:
//
// * "S": second
//
// * "M": minute
//
// * "H": hour
//
// * "D": day
//
// Examples:
//
// * 5 reqs/second: "5-S"
//
// * 10 reqs/minute: "10-M"
//
// * 1000 reqs/hour: "1000-H"
//
// * 2000 reqs/day: "2000-D"
func WithFormattedLimit(v string) ServiceInitOption {
	return func(o *Service) {
		o.formattedLimit = v
	}
}

func WithPeriodLimit(period time.Duration, limit uint64) ServiceInitOption {
	return func(o *Service) {
		o.period = period
		o.limit = limit
	}
}

/* ServiceMethodOption */
type ServiceMethodOption func(*serviceOptions)

type serviceOptions struct {
	cacheKey string
}

func WithCacheKey(v string) ServiceMethodOption {
	return func(o *serviceOptions) {
		o.cacheKey = v
	}
}
