package app

import (
	"fmt"
	"net"
	"service-user/database"
	"service-user/dto"
	defaultCtrl "service-user/module/default/controller"
	userCtrl "service-user/module/user/controller"

	"service-user/module/user/repository"
	"service-user/module/user/usecase"

	"google.golang.org/grpc"

	"log"
	"service-user/pkg"
	pb_user "service-user/proto/user"
)

func init() {
	pkg.LoadConfig(".env")

	//min := 3000
	//max := 3010
	//dto.CfgApp.RestPort = rand.Intn(max-min) + min
	//dto.CfgApp.GRPCPort = rand.Intn(max-min) + (min + 1)
	dto.CfgApp.GRPCPort = 3006
}

func NewGRPC() error {
	pkg.NewConsul(dto.CfgApp.ServiceName, dto.CfgApp.GRPCPort, "GRPC")

	db := database.SetupDatabase()
	userRepository := repository.NewUserRepository(db)
	userService := usecase.NewUserUsecase(userRepository)

	InitUser := userCtrl.NewHandlerRPCUser(userService)
	InitHealth := userCtrl.NewhealthCheck()

	s := grpc.NewServer()
	pb_user.RegisterHealthServer(s, InitHealth)
	pb_user.RegisterServiceUserRPCServer(s, InitUser)

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

	userRepo := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUsecase(userRepo)

	userCtrl.NewUserControllerHTTP(httpServer, userUseCase)
	defaultCtrl.InitDefaultController(httpServer)

	//pkg.NewConsul(dto.CfgApp.ServiceName, dto.CfgApp.RestPort, "REST")

	log.Println("Starting Rest API server at", dto.CfgApp.RestPort)
	err := httpServer.Run(fmt.Sprintf(`:%v`, dto.CfgApp.RestPort))
	if err != nil {
		panic(err)
	}

}
