package natsconn

type Config struct {
	Enabled  bool   `yaml:"enabled" validate:"required,boolean" default:"true"`
	Addr     string `validate:"required,hostname_port" example:"localhost:4222"`
	User     string `secret:"true"`
	Password string `secret:"true"`
}
