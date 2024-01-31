package dto

type ConfigEnvironment struct {
	Database ConfigDatabase
	App      ConfigApp
	Consul   ConfigConsul
	Kafka    ConfigKafka
}

var CfgConsul ConfigConsul
var CfgApp ConfigApp
var CfgDB ConfigDatabase
var CfgKafka ConfigKafka

type ConfigApp struct {
	ServiceName string `env:"SERVICE_NAME"`
	Timezone    string `env:"TIMEZONE"`
	Version     string `env:"VERSION"`
	RestPort    int    `env:"REST_PORT"`
	GRPCPort    int    `env:"GRPC_PORT"`
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
	ConsulHost string `env:"CONSUL_HOST"`
	ConsulPort int    `env:"CONSUL_PORT"`
}

type ConfigKafka struct {
	KafkaUser     string `env:"KAFKA_USER"`
	KafkaPassword string `env:"KAFKA_PASSWORD"`
	KafkaAddress  string `env:"KAFKA_ADDRESS"`
}
