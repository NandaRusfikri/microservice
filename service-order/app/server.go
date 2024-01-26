package app

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	hv1 "google.golang.org/grpc/health/grpc_health_v1"
	"math/rand"
	"net"
	"service-order/database"
	"service-order/dto"
	orderCtrl "service-order/module/order/controller"
	order_repo "service-order/module/order/repository"
	order_serv "service-order/module/order/service"
	product_repo "service-order/module/product/repository"
	"service-order/pkg"
	"service-order/pkg/kafka"
	pb_order "service-order/proto/order"
)

func init() {
	pkg.LoadConfig(".env")

	min := 3100
	max := 3150
	dto.CfgApp.RestPort = rand.Intn(max-min) + min
	dto.CfgApp.GRPCPort = 6000
}

func NewGRPC() error {
	pkg.NewConsul(dto.CfgApp.ServiceName, dto.CfgApp.GRPCPort, "GRPC")

	kafka := kafka.NewKafka()

	defer func() {
		if err := kafka.Producer.Close(); err != nil {
			log.Errorf("Unable to stop kafka producer: %v", err)
			return
		}
	}()

	db := database.SetupDatabase()
	orderRepo := order_repo.NewOrderRepositorySQL(db)
	productRepo := product_repo.NewOrderRepositoryGRPC()
	orderService := order_serv.NewOrderService(orderRepo, productRepo, kafka)

	InitServiceGRPC := orderCtrl.NewOrderControllerGRPC(orderService)

	s := grpc.NewServer()
	healthServer := health.NewServer()
	healthServer.SetServingStatus("", hv1.HealthCheckResponse_SERVING)
	hv1.RegisterHealthServer(s, health.NewServer())
	pb_order.RegisterServiceOrderRPCServer(s, InitServiceGRPC)

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
	//kafka := pkg.NewKafka()
	//
	//orderRepo := order_repo.NewOrderRepositorySQL(db)
	//productRepo := product_repo.NewOrderRepositoryGRPC()
	//orderService := order_serv.NewOrderService(orderRepo, productRepo, kafka)
	//
	//orderCtrl.NewOrderControllerRestAPI(orderService, httpServer)
	//defaultCtrl.InitDefaultController(httpServer)
	//
	//pkg.NewConsul(dto.CfgApp.ServiceName, dto.CfgApp.GRPCPort, "REST")
	//log.Println("Starting REST server at", dto.CfgApp.RestPort)
	//err := httpServer.Run(fmt.Sprintf(`:%v`, dto.CfgApp.RestPort))
	//if err != nil {
	//	panic(err)
	//}

}
