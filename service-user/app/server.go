package app

import (
	"fmt"
	"math/rand"
	"service-user/database"
	"service-user/dto"
	"time"

	"net"
	defaultCtrl "service-user/module/default/controller"
	userCtrl "service-user/module/user/controller"

	"service-user/module/user/repository"
	"service-user/module/user/usecase"

	"google.golang.org/grpc"

	"log"
	"service-user/pkg"
	pb_user "service-user/proto/user"
)

func NewGRPC() error {
	pkg.NewConsul(dto.CfgApp.ServiceName, dto.CfgApp.GRPCPort)

	db := database.SetupDatabase()
	userRepository := repository.NewUserRepository(db)
	userService := usecase.NewUserService(userRepository)

	InitUser := userCtrl.NewHandlerRPCUser(userService)

	s := grpc.NewServer()
	pb_user.RegisterServiceUserRPCServer(s, InitUser)

	log.Println("Starting RPC server at", dto.CfgApp.GRPCPort)
	l, err := net.Listen("tcp", fmt.Sprintf(":%v", dto.CfgApp.GRPCPort))
	if err != nil {
		log.Fatalf("could not listen to %v: %v", dto.CfgApp.GRPCPort, err)
	}

	return s.Serve(l)
}

func init() {
	pkg.LoadConfig(".env")
	rand.NewSource(time.Now().UnixNano())
	dto.CfgApp.RestPort = rand.Intn(4001) + 1000
	fmt.Println("dto.CfgApp.RestPort ", dto.CfgApp.RestPort)
	rand.NewSource(time.Now().UnixNano())
	dto.CfgApp.GRPCPort = rand.Intn(4001) + 1000
	fmt.Println("dto.CfgApp.GRPCPort ", dto.CfgApp.GRPCPort)
}

func NewRestAPI() {

	db := database.SetupDatabase()
	httpServer := pkg.InitHTTPGin()

	userRepo := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserService(userRepo)

	userCtrl.NewUserControllerHTTP(httpServer, userUseCase)
	defaultCtrl.InitDefaultController(httpServer)

	err := httpServer.Run(fmt.Sprintf(`:%v`, dto.CfgApp.RestPort))
	if err != nil {
		panic(err)
	}

}
