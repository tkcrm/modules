package dragonfly

type Config struct {
	Addr         string `validate:"required,hostname_port" example:"localhost:6379"`
	User         string `secret:"true"`
	Password     string `secret:"true"`
	DbIndex      int    `validate:"gte=0" default:"0"`
	PingInterval int    `validate:"gt=0" default:"10"`
}
