package postgres

type Config struct {
	Enabled      bool   `yaml:"enabled" validate:"required,boolean" default:"true"`
	Addr         string `validate:"required,hostname_port" example:"localhost:5432"`
	User         string `validate:"required" secret:"true"`
	Password     string `validate:"required" secret:"true"`
	DBName       string `yaml:"db_name" validate:"required" env:"DB_NAME"`
	PingInterval int    `yaml:"ping_interval" validate:"required,gt=0" default:"10"`
	MinConns     int32  `yaml:"min_conns" validate:"required,gt=0" default:"3"`
	MaxConns     int32  `yaml:"max_conns" validate:"required,gt=0" default:"6"`
}
