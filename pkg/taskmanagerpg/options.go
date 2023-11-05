package taskmanagerpg

import "time"

type Option func(*Options)

type Options struct {
	RunAt       time.Time
	QueueName   string
	Priority    int16
	JSONPayload bool
}

func WithRunAt(v time.Time) Option {
	return func(o *Options) {
		o.RunAt = v
	}
}

func WithQueueName(v string) Option {
	return func(o *Options) {
		o.QueueName = v
	}
}

func WithPriority(v int16) Option {
	return func(o *Options) {
		o.Priority = v
	}
}

func WithJSONPayload() Option {
	return func(o *Options) {
		o.JSONPayload = true
	}
}
