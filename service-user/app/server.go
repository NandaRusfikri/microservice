package app

import (
	"context"
	"flag"
	"fmt"
	"net"
	"service-user/constant"
	"service-user/database"
	"service-user/dto"
	defaultCtrl "service-user/module/default/controller"
	userCtrl "service-user/module/user/controller"
	"service-user/pkg/kafka"

	"service-user/module/user/repository"
	"service-user/module/user/usecase"

	"google.golang.org/grpc"

	log "github.com/sirupsen/logrus"
	"service-user/pkg"
	pb_user "service-user/proto/user"
)

func init() {
	pkg.LoadConfig(".env")
	portGRPC := flag.Int("grpc-port", 4100, "GRPC Port Service")
	portREST := flag.Int("rest-port", 4200, "REST Port Service")
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

	db := database.SetupDatabase()
	userRepository := repository.NewUserRepository(db)
	userService := usecase.NewUserUsecase(userRepository, kafkaProducer)

	go kafka.NewKafkaConsumer(ctx, []string{constant.TOPIC_NEW_ORDER})
	userCtrl.NewUserControllerKafka(userService)

	s := grpc.NewServer()
	pb_user.RegisterHealthServer(s, defaultCtrl.NewhealthCheck())
	pb_user.RegisterServiceUserRPCServer(s, userCtrl.NewHandlerRPCUser(userService))

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
	//userRepo := repository.NewUserRepository(db)
	//userUseCase := usecase.NewUserUsecase(userRepo)
	//
	//userCtrl.NewUserControllerHTTP(httpServer, userUseCase)
	//defaultCtrl.InitDefaultController(httpServer)
	//
	////pkg.NewConsul(dto.CfgApp.ServiceName, dto.CfgApp.RestPort, "REST")
	//
	//log.Println("Starting Rest API server at", dto.CfgApp.RestPort)
	//err := httpServer.Run(fmt.Sprintf(`:%v`, dto.CfgApp.RestPort))
	//if err != nil {
	//	panic(err)
	//}

}
