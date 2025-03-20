package redis

type Config struct {
	Enabled      bool   `yaml:"enabled" validate:"required,boolean" default:"true"`
	Addr         string `validate:"required,hostname_port" example:"localhost:6379"`
	User         string `secret:"true"`
	Password     string `secret:"true"`
	DbIndex      int    `yaml:"db_index" validate:"gte=0" default:"0"`
	PingInterval int    `yaml:"ping_interval" validate:"gt=0" default:"10"`
}
