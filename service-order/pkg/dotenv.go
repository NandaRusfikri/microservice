package pkg

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
	"service-order/constant"
	"service-order/dto"
	"strconv"
)

func LoadConfig(path string) {
	err := godotenv.Load(path)
	if err != nil {
		log.Errorln("Error loading .env file ", err)
	}

	dto.CfgDB = dto.ConfigDatabase{
		DBHost:    os.Getenv("DB_HOST"),
		DBPort:    os.Getenv("DB_PORT"),
		DBUser:    os.Getenv("DB_USER"),
		DBName:    os.Getenv("DB_NAME"),
		DBPass:    os.Getenv("DB_PASS"),
		DBSSLmode: os.Getenv("DB_SSLMODE"),
	}

	dto.CfgApp = dto.ConfigApp{
		ServiceName: constant.SERVICE_NAME,
		Timezone:    os.Getenv("TIMEZONE"),
		Version:     os.Getenv("VERSION"),
		SwaggerHost: os.Getenv("SWAGGER_HOST"),
	}
	dto.CfgKafka = dto.ConfigKafka{
		KafkaAddress:  os.Getenv("KAFKA_ADDRESS"),
		KafkaUser:     os.Getenv("KAFKA_USER"),
		KafkaPassword: os.Getenv("KAFKA_PASSWORD"),
	}
	dto.CfgApp.RestPort, _ = strconv.Atoi(os.Getenv("REST_PORT"))

	dto.CfgConsul.ConsulHost = os.Getenv("CONSUL_HOST")
	dto.CfgConsul.ConsulPort, _ = strconv.Atoi(os.Getenv("CONSUL_PORT"))

}
