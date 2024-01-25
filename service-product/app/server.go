package app

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	hv1 "google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"math/rand"
	"net"
	"service-product/database"
	"service-product/dto"
	defaultCtrl "service-product/module/default/controller"
	userCtrl "service-product/module/product/controller"
	"service-product/module/product/repository"
	"service-product/module/product/usecase"
	"service-product/pkg"
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

	db := database.SetupDatabase()
	userRepository := repository.NewRepository(db)
	userService := usecase.NewServiceProduct(userRepository)

	InitUser := userCtrl.NewControllerProductRPC(userService)

	s := grpc.NewServer()
	healthServer := health.NewServer()
	healthServer.SetServingStatus("", hv1.HealthCheckResponse_SERVING)
	hv1.RegisterHealthServer(s, health.NewServer())
	pb_user.RegisterServiceProductRPCServer(s, InitUser)

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

	userCtrl.NewControllerProductHTTP(httpServer, userUseCase)
	defaultCtrl.InitDefaultController(httpServer)

	pkg.NewConsul(dto.CfgApp.ServiceName, dto.CfgApp.GRPCPort, "REST")

	log.Println("Starting REST server at", dto.CfgApp.RestPort)
	err := httpServer.Run(fmt.Sprintf(`:%v`, dto.CfgApp.RestPort))
	if err != nil {
		panic(err)
	}

}
