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
	UniqueTasks  bool         `yaml:"unique_tasks"`
	RedisConfig  RedisConfig  `yaml:"redis_config"`
	ServerConfig ServerConfig `yaml:"server_config"`
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
