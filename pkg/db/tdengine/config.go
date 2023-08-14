package tdengine

type Config struct {
	Addr     string `validate:"required,hostname_port" example:"localhost:6030"`
	User     string `validate:"required" secret:"true"`
	Password string `validate:"required" secret:"true"`
	DBName   string `validate:"required" env:"DB_NAME"`
	// In seconds. Default 10 seconds
	PingInterval int `validate:"gt=0" default:"10"`
}
