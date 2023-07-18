package taskmanager

import "time"

type Option func(*options)

type options struct {
	runAt       time.Time
	queueName   string
	priority    int16
	jsonPayload bool
}

func WithRunAt(v time.Time) Option {
	return func(o *options) {
		o.runAt = v
	}
}

func WithQueueName(v string) Option {
	return func(o *options) {
		o.queueName = v
	}
}

func WithPriority(v int16) Option {
	return func(o *options) {
		o.priority = v
	}
}

func WithJSONPayload(v bool) Option {
	return func(o *options) {
		o.jsonPayload = v
	}
}
