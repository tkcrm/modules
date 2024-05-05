package bunconn

type Config struct {
	Addr     string `validate:"required,hostname_port" example:"localhost:5432"`
	User     string `validate:"required" secret:"true"`
	Password string `validate:"required" secret:"true"`
	DBName   string `validate:"required" env:"DB_NAME" yaml:"db_name"`
	BUNDEBUG bool   `env:"BUNDEBUG"`
	// In seconds. Default 10 seconds
	PingInterval int `validate:"required,gt=0" default:"10" usage:"in seconds" yaml:"ping_interval"`
}
