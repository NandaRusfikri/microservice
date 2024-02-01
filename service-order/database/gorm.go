package database

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"service-order/dto"
	entity "service-order/module/order/entity"
)

func SetupDatabase() *gorm.DB {
	log.Debug("🔌 Connecting into Database")
	dbHost := dto.CfgDB.DBHost
	dbUsername := dto.CfgDB.DBUser
	dbPassword := dto.CfgDB.DBPass
	dbName := dto.CfgDB.DBName
	dbPort := dto.CfgDB.DBPort
	dbSSLMode := dto.CfgDB.DBSSLmode

	path := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v",
		dbHost, dbUsername, dbPassword, dbName, dbPort, dbSSLMode, dto.CfgApp.Timezone)

	db, err := gorm.Open(postgres.Open(path), &gorm.Config{})

	if err != nil {
		defer log.Info("Connect into Database Failed")
		log.Fatal(err.Error())
	}

	err = db.AutoMigrate(
		&entity.Order{},
		&entity.TopicOrderReply{},
	)

	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}
