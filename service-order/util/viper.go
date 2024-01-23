package util

import (
	"github.com/gin-contrib/cors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"service-order/constant"
	"service-order/dto"
	"strings"
	"time"
)

func InitConfig() {
	log.Info("InitConfig() - starting...")

	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("InitConfig() - error: ", err)
	}

	if viper.GetBool(constant.CONFIG_PATH_APP_DEBUG) {
		log.Info("Service run on DEBUG mode.")
	}
	log.Info("InitConfig() - finished.")
}
func InitConfigTest() {
	log.Info("InitConfigTest() - starting...")

	//if *isDevelopment {
	//	viper.SetConfigFile(constant.CONFIG_PATH_DEVELOPMENT)
	//} else {
	//	viper.SetConfigFile(constant.CONFIG_PATH)
	//}
	viper.SetConfigFile(constant.CONFIG_PATH_DEVELOPMENT)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("InitConfigTest() - error: ", err)
	}

	if viper.GetBool(constant.CONFIG_PATH_APP_DEBUG) {
		log.Info("Service run on DEBUG mode.")
	}
	log.Info("InitConfigTest() - finished.")
}
func GetDatabaseConfig() *dto.DatabaseConfig {
	log.Info("GetDatabaseConfig() - starting...")
	dbHost := viper.GetString(constant.CONFIG_PATH_DB_HOST)
	dbPort := viper.GetString(constant.CONFIG_PATH_DB_PORT)
	dbUsername := viper.GetString(constant.CONFIG_PATH_DB_USERNAME)
	dbPassword := viper.GetString(constant.CONFIG_PATH_DB_PASSWORD)
	dbName := viper.GetString(constant.CONFIG_PATH_DB_DATABASE)
	dbDialect := viper.GetString(constant.CONFIG_PATH_DB_DIALECT)
	dbSSLMode := viper.GetString(constant.CONFIG_PATH_DB_SSL_MODE)
	dbTimeZone := viper.GetString(constant.CONFIG_PATH_APP_TIME_ZONE)

	if dbHost == "" {
		dbHost = constant.DEFAULT_DB_HOST
	}

	if dbPort == "" {
		dbPort = constant.DEFAULT_DB_PORT
	}

	if dbUsername == "" {
		dbUsername = constant.DEFAULT_DB_USERNAME
	}

	if dbPassword == "" {
		dbPassword = constant.DEFAULT_DB_PASSWORD
	}

	if dbName == "" {
		dbName = constant.DEFAULT_DB_DATABASE
	}

	if dbDialect == "" {
		dbDialect = constant.DEFAULT_DB_DIALECT
	}

	if dbSSLMode == "" {
		dbSSLMode = constant.DEFAULT_DB_SSL_MODE
	}

	if dbTimeZone == "" {
		dbTimeZone = constant.DEFAULT_TIME_ZONE
	}

	config := &dto.DatabaseConfig{
		Host:     dbHost,
		Port:     dbPort,
		Username: dbUsername,
		Password: dbPassword,
		Database: dbName,
		Dialect:  dbDialect,
		SSLMode:  dbSSLMode,
		TimeZone: dbTimeZone,
	}

	log.Info("GetDatabaseConfig() - finished.")
	return config
}

func GetAppConfig() *dto.AppConfig {
	log.Debug("GetAppConfig() - starting...")
	appDebug := viper.GetBool(constant.CONFIG_PATH_APP_DEBUG)
	appHost := viper.GetString(constant.CONFIG_PATH_APP_HOST)
	appPort := viper.GetString(constant.CONFIG_PATH_APP_PORT)

	appUseTLS := viper.GetBool(constant.CONFIG_PATH_APP_USE_TLS)
	appCertFilePath := viper.GetString(constant.CONFIG_PATH_APP_CERT_FILE_PATH)
	appKeyFilePath := viper.GetString(constant.CONFIG_PATH_APP_KEY_FILE_PATH)

	appCors := cors.DefaultConfig()
	appCors.AllowAllOrigins = true

	if !appDebug {
		corsAllowOriginS := viper.GetString(constant.CONFIG_PATH_ALLOW_ORIGIN)
		corsAllowOrigin := strings.Split(corsAllowOriginS, ",")
		corsAllowMethodS := viper.GetString(constant.CONFIG_PATH_ALLOW_METHOD)
		corsAllowMethod := strings.Split(corsAllowMethodS, ",")
		corsAllowHeaderS := viper.GetString(constant.CONFIG_PATH_ALLOW_HEADER)
		corsAllowHeader := strings.Split(corsAllowHeaderS, ",")
		corsExposeHeaderS := viper.GetString(constant.CONFIG_PATH_EXPOSE_HEADER)
		corsExposeHeader := strings.Split(corsExposeHeaderS, ",")
		corsAllowCredential := viper.GetBool(constant.CONFIG_PATH_ALLOW_CREDENTIAL)
		corsAllowWildcard := viper.GetBool(constant.CONFIG_PATH_ALLOW_WILDCARD)
		corsAllowBrowserExtension := viper.GetBool(constant.CONFIG_PATH_ALLOW_BROWSER_EXTENSION)
		corsAllowWebSocket := viper.GetBool(constant.CONFIG_PATH_ALLOW_WEB_SOCKET)
		corsAllowFile := viper.GetBool(constant.CONFIG_PATH_ALLOW_FILE)
		maxAgeI := viper.GetDuration(constant.CONFIG_PATH_MAX_AGE)
		maxAge := maxAgeI * time.Hour

		appCors = cors.Config{
			AllowAllOrigins:        false,
			AllowOrigins:           corsAllowOrigin,
			AllowMethods:           corsAllowMethod,
			AllowHeaders:           corsAllowHeader,
			ExposeHeaders:          corsExposeHeader,
			AllowCredentials:       corsAllowCredential,
			AllowBrowserExtensions: corsAllowBrowserExtension,
			AllowWildcard:          corsAllowWildcard,
			AllowWebSockets:        corsAllowWebSocket,
			AllowFiles:             corsAllowFile,
			MaxAge:                 maxAge,
		}
	}

	if appHost == "" {
		appHost = constant.DEFAULT_APP_HOST
	}

	if appPort == "" {
		appPort = constant.DEFAULT_APP_PORT
	}

	log.Debug("GetAppConfig() - finished.")

	return &dto.AppConfig{
		Debug:        appDebug,
		Host:         appHost,
		Port:         appPort,
		UseTLS:       appUseTLS,
		CertFilePath: appCertFilePath,
		KeyFilePath:  appKeyFilePath,
		Cors:         &appCors,
	}
}




