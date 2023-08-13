package natsconn

type Config struct {
	Addr     string `validate:"required,hostname_port" example:"localhost:5432"`
	User     string
	Password string `secret:"true"`
}
