package app

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	hv1 "google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"math/rand"
	"net"
	"service-order/database"
	"service-order/dto"
	defaultCtrl "service-order/module/default/controller"
	orderCtrl "service-order/module/order/controller"
	order_repo "service-order/module/order/repository"
	order_serv "service-order/module/order/service"
	product_repo "service-order/module/product/repository"
	"service-order/pkg"
	pb_order "service-order/proto/order"
	"time"
)

func init() {
	pkg.LoadConfig(".env")
	rand.NewSource(time.Now().UnixNano())
	dto.CfgApp.RestPort = rand.Intn(10) + 5000
	dto.CfgApp.GRPCPort = rand.Intn(10+1) + 5000
	fmt.Println("dto.CfgApp.GRPCPort ", dto.CfgApp.GRPCPort)
	fmt.Println("dto.CfgApp.RestPort ", dto.CfgApp.RestPort)
}

func NewGRPC() error {
	pkg.NewConsul(dto.CfgApp.ServiceName, dto.CfgApp.GRPCPort)

	db := database.SetupDatabase()
	orderRepo := order_repo.NewOrderRepositorySQL(db)
	productRepo := product_repo.NewOrderRepositoryGRPC()
	orderService := order_serv.NewOrderService(orderRepo, productRepo)

	InitServiceGRPC := orderCtrl.NewOrderControllerGRPC(orderService)

	s := grpc.NewServer()
	healthServer := health.NewServer()
	healthServer.SetServingStatus("", hv1.HealthCheckResponse_SERVING)
	hv1.RegisterHealthServer(s, health.NewServer())
	pb_order.RegisterServiceOrderRPCServer(s, InitServiceGRPC)

	log.Println("Starting RPC server at", dto.CfgApp.GRPCPort)
	l, err := net.Listen("tcp", fmt.Sprintf(":%v", dto.CfgApp.GRPCPort))
	if err != nil {
		log.Fatalf("could not listen to %v: %v", dto.CfgApp.GRPCPort, err)
	}

	return s.Serve(l)
}

func NewRestAPI() {

	db := database.SetupDatabase()
	httpServer := pkg.InitHTTPGin()

	orderRepo := order_repo.NewOrderRepositorySQL(db)
	productRepo := product_repo.NewOrderRepositoryGRPC()
	orderService := order_serv.NewOrderService(orderRepo, productRepo)

	orderCtrl.NewOrderControllerRestAPI(orderService, httpServer)
	defaultCtrl.InitDefaultController(httpServer)

	err := httpServer.Run(fmt.Sprintf(`:%v`, dto.CfgApp.RestPort))
	if err != nil {
		panic(err)
	}

}
