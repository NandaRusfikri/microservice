package app

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"service-product/constant"

	log "github.com/sirupsen/logrus"
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

	portGRPC := flag.Int("grpc-port", 5100, "GRPC Port Service")
	portREST := flag.Int("rest-port", 5200, "REST Port Service")
	flag.Parse()

	dto.CfgApp.RestPort = *portREST
	dto.CfgApp.GRPCPort = *portGRPC
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

	go kafka.NewKafkaConsumer(ctx, []string{constant.TOPIC_NEW_ORDER})

	db := database.SetupDatabase()
	productRepository := repository.NewRepository(db)
	productService := usecase.NewServiceProduct(productRepository, kafkaProducer)

	InitProductGRPC := productCtrl.NewControllerProductGRPC(productService)
	productCtrl.NewProductControllerKafka(productService)

	InitHealth := defaultCtrl.NewhealthCheck()

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

	//db := database.SetupDatabase()
	//httpServer := pkg.InitHTTPGin()
	//
	//userRepo := repository.NewRepository(db)
	//userUseCase := usecase.NewServiceProduct(userRepo)
	//
	//productCtrl.NewControllerProductHTTP(httpServer, userUseCase)
	//defaultCtrl.InitDefaultController(httpServer)
	//
	//pkg.NewConsul(dto.CfgApp.ServiceName, dto.CfgApp.GRPCPort, "REST")
	//
	//log.Println("Starting REST server at", dto.CfgApp.RestPort)
	//err := httpServer.Run(fmt.Sprintf(`:%v`, dto.CfgApp.RestPort))
	//if err != nil {
	//	panic(err)
	//}

}
