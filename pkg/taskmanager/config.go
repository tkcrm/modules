package taskmanager

import (
	"time"

	"github.com/hibiken/asynq"
)

const (
	defaultName                         = "taskmanager"
	defaultGracefulShutdownTimeDuration = time.Second * 20
)

type Config struct {
	Name         string
	UniqueTasks  bool
	RedisConfig  RedisConfig
	ServerConfig ServerConfig
}

type RedisConfig struct {
	Addr     string
	Username string
	Password string
	DB       int
}

type ServerConfig struct {
	asynq.Config
}
