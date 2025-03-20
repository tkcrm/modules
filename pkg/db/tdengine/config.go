package tdengine

type Config struct {
	Enabled  bool   `yaml:"enabled" validate:"required,boolean" default:"true"`
	Addr     string `validate:"required,hostname_port" example:"localhost:6030"`
	User     string `validate:"required" secret:"true"`
	Password string `validate:"required" secret:"true"`
	DBName   string `yaml:"db_name" validate:"required" env:"DB_NAME"`
	// In seconds. Default 10 seconds
	PingInterval int `yaml:"ping_interval" validate:"gt=0" default:"10"`
}
