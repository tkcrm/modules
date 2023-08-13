package postgres

type Config struct {
	Addr         string `validate:"required,hostname_port" example:"localhost:5432"`
	User         string `validate:"required"`
	Password     string `validate:"required" secret:"true"`
	DBName       string `validate:"required" env:"DB_NAME"`
	PingInterval int    `validate:"required,gt=0" default:"10"`
	MinConns     int32  `validate:"required,gt=0" default:"3"`
	MaxConns     int32  `validate:"required,gt=0" default:"6"`
}
