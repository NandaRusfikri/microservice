package app

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"service-product/constant"

	log "github.com/sirupsen/logrus"
	"math/rand"
	"net"
	"service-product/database"
	"service-product/dto"
	defaultCtrl "service-product/module/default/controller"
	productCtrl "service-product/module/product/controller"
	"service-product/module/product/repository"
	"service-product/module/product/usecase"
	"service-product/pkg"
	"service-product/pkg/kafka"
	pb_health "service-product/proto/health"
	pb_user "service-product/proto/product"
)

func init() {
	pkg.LoadConfig(".env")

	min := 4000
	max := 4100
	dto.CfgApp.RestPort = rand.Intn(max-min) + min
	dto.CfgApp.GRPCPort = rand.Intn(max-min) + (min + 1)
}

func NewGRPC() error {
	pkg.NewConsul(dto.CfgApp.ServiceName, dto.CfgApp.GRPCPort, "GRPC")

	kafkaProducer := kafka.NewKafkaProducer()

	defer func() {
		if err := kafkaProducer.Producer.Close(); err != nil {
			log.Errorf("Unable to stop kafka producer: %v", err)
			return
		}
	}()

	ctx := context.Background()

	go kafka.NewKafkaConsumer(ctx, []string{constant.TOPIC_PRODUCT_STOCK_UPDATE})

	pkg.NewConsul(dto.CfgApp.ServiceName, dto.CfgApp.GRPCPort, "GRPC")

	db := database.SetupDatabase()
	productRepository := repository.NewRepository(db)
	productService := usecase.NewServiceProduct(productRepository)

	InitProductGRPC := productCtrl.NewControllerProductGRPC(productService)
	productCtrl.NewProductControllerKafka(productService)

	InitHealth := productCtrl.NewhealthCheck()

	s := grpc.NewServer()
	pb_health.RegisterHealthServer(s, InitHealth)

	pb_user.RegisterServiceProductRPCServer(s, InitProductGRPC)

	log.Println("Starting GRPC server at", dto.CfgApp.GRPCPort)
	l, err := net.Listen("tcp", fmt.Sprintf(":%v", dto.CfgApp.GRPCPort))
	if err != nil {
		log.Fatalf("could not listen to %v: %v", dto.CfgApp.GRPCPort, err)
	}

	return s.Serve(l)
}

func NewRestAPI() {

	db := database.SetupDatabase()
	httpServer := pkg.InitHTTPGin()

	userRepo := repository.NewRepository(db)
	userUseCase := usecase.NewServiceProduct(userRepo)

	productCtrl.NewControllerProductHTTP(httpServer, userUseCase)
	defaultCtrl.InitDefaultController(httpServer)

	pkg.NewConsul(dto.CfgApp.ServiceName, dto.CfgApp.GRPCPort, "REST")

	log.Println("Starting REST server at", dto.CfgApp.RestPort)
	err := httpServer.Run(fmt.Sprintf(`:%v`, dto.CfgApp.RestPort))
	if err != nil {
		panic(err)
	}

}
