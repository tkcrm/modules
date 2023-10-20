package taskmanager

import "time"

const (
	defaultQueueName = "default"
)

type Option func(*options)

type options struct {
	jsonPayload bool
	retry       int
	queue       string
	taskID      string
	timeout     time.Duration
	deadline    time.Time
	uniqueTTL   time.Duration
	processAt   time.Time
	retention   time.Duration
	group       string
}

func WithJSONPayload(v bool) Option {
	return func(o *options) {
		o.jsonPayload = v
	}
}

func WithMaxRetry(v int) Option {
	return func(o *options) {
		o.retry = v
	}
}

func WithQueue(v string) Option {
	return func(o *options) {
		o.queue = v
	}
}

func WithTaskID(v string) Option {
	return func(o *options) {
		o.taskID = v
	}
}

func WithTimeout(v time.Duration) Option {
	return func(o *options) {
		o.timeout = v
	}
}

func WithDeadline(v time.Time) Option {
	return func(o *options) {
		o.deadline = v
	}
}

func WithUniqueTTL(v time.Duration) Option {
	return func(o *options) {
		o.uniqueTTL = v
	}
}

func WithProcessAt(v time.Time) Option {
	return func(o *options) {
		o.processAt = v
	}
}

func WithRetention(v time.Duration) Option {
	return func(o *options) {
		o.retention = v
	}
}

func WithGroup(v string) Option {
	return func(o *options) {
		o.group = v
	}
}
