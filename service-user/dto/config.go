package dto

type ConfigEnvironment struct {
	Database ConfigDatabase
	App      ConfigApp
	Consul   ConfigConsul
}

var CfgConsul ConfigConsul
var CfgApp ConfigApp
var CfgDB ConfigDatabase

type ConfigApp struct {
	ServiceName string `env:"SERVICE_NAME"`
	Timezone    string `env:"TIMEZONE"`
	Version     string `env:"VERSION"`
	RestPort    int    `env:"REST_PORT"`
	SwaggerHost string `env:"SWAGGER_HOST"`
}
type ConfigDatabase struct {
	DBUser    string `env:"DB_USER"`
	DBPass    string `env:"DB_PASS"`
	DBHost    string `env:"DB_HOST"`
	DBPort    string `env:"DB_PORT"`
	DBName    string `env:"DB_NAME"`
	DBSSLmode string `env:"DB_SSLMODE"`
}

type ConfigConsul struct {
	ServiceName string `env:"SERVICE_NAME"`
	ServicePort int    `env:"SERVICE_PORT"`
	ConsulHost  string `env:"CONSUL_HOST"`
	ConsulPort  int    `env:"CONSUL_PORT"`
}
