package mqttconn

type Config struct {
	Addr     string `validate:"required,hostname_port" example:"localhost:1883"`
	User     string
	Password string `secret:"true"`
}
