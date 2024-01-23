package main

import (
	"fmt"
	PluginConsul "github.com/go-micro/plugins/v4/registry/consul"
	ConsulAPI "github.com/hashicorp/consul/api"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"service-order/entities"
	"service-order/pkg"

	handlers "service-order/controllers/order"

	pb_order "service-order/proto/order"
	repositorys "service-order/repositorys/order"
	services "service-order/services/order"


	"go-micro.dev/v4"

	"github.com/go-micro/plugins/v4/server/grpc"

)

func main() {

	db := SetupDatabase()
	orderRepository := repositorys.NewOrderRepositorySQL(db)
	orderRepositoryRPC := repositorys.NewOrderRepositoryGRPC()
	orderService := services.NewOrderService(orderRepository,orderRepositoryRPC)

	 InitOrderController := handlers.NewHandlerRPCOrder(orderService)

	 //micro.Client()

	service := micro.NewService(
		micro.Server(grpc.NewServer()),
		micro.Name("service-order"),
	)
	consulConf := ConsulAPI.DefaultConfig()
	consulConf.Address = "127.0.0.1:8567"
	consulConf.Scheme = "http"


	service.Init(
		micro.Registry(
			PluginConsul.NewRegistry(
				PluginConsul.Config(consulConf))))

	fmt.Println()


	pb_order.RegisterServiceOrderHandlerRPCHandler(service.Server(), InitOrderController)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
func SetupDatabase() *gorm.DB {
	urldb := pkg.GodotEnv("DATABASE_URI")
	//fmt.Println("urldb ",urldb)
	db, err := gorm.Open(mysql.Open(urldb), &gorm.Config{})

	if err != nil {
		defer log.Info("Connect into Database Failed")
		log.Fatal(err.Error())
	}

	if os.Getenv("GO_ENV") != "production" {
		log.Info("Connect into Database Successfully")
	}

	err = db.AutoMigrate(
		&entities.EntityOrder{},
		//&models.ModelUser{},
	)

	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}


