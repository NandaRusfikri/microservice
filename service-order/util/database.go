package util

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"service-order/entities"
	"service-order/pkg"
)

func SetupDatabase() *gorm.DB {
	urldb := pkg.GodotEnv("DATABASE_URI")
	//fmt.Println("urldb ",urldb)
	db, err := gorm.Open(mysql.Open(urldb), &gorm.Config{})

	if err != nil {
		defer logrus.Info("Connect into Database Failed")
		logrus.Fatal(err.Error())
	}

	if os.Getenv("GO_ENV") != "production" {
		logrus.Info("Connect into Database Successfully")
	}

	err = db.AutoMigrate(
		&entities.EntityOrder{},
		//&models.ModelUser{},
	)

	if err != nil {
		logrus.Fatal(err.Error())
	}

	return db
}
