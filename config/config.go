package config

// AppConfig - Application Configurations
type AppConfig struct {
	GrpcServerPort int    `env:"GRPC_SERVER_PORT,default=50060"`
	Dbusername     string `env:"DB_USER,default=root"`
	Dbpassword     string `env:"DB_PASS,default=password"`
	Dbdatabase     string `env:"DB_DATABASE,default=gorpc"`
	Dbport         int    `env:"DB_PORT,default=5432"`
}
