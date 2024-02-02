package app

import (
	"context"
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"service-order/constant"
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

	portGRPC := flag.Int("grpc-port", 6100, "GRPC Port Service")
	portREST := flag.Int("rest-port", 6200, "REST Port Service")
	flag.Parse()

	dto.CfgApp.RestPort = *portREST
	dto.CfgApp.GRPCPort = *portGRPC

	pkg.NewConsul(dto.CfgApp.ServiceName, dto.CfgApp.GRPCPort, "GRPC")
	pkg.NewConsul(dto.CfgApp.ServiceName, dto.CfgApp.RestPort, "REST")
}

func NewGRPC() error {

	kafkaProducer := kafka.NewKafkaProducer()

	defer func() {
		if err := kafkaProducer.Producer.Close(); err != nil {
			log.Errorf("Unable to stop kafka producer: %v", err)
			return
		}
	}()

	ctx := context.Background()

	go kafka.NewKafkaConsumer(ctx, []string{constant.TOPIC_NEW_ORDER, constant.TOPIC_ORDER_REPLY})

	db := database.SetupDatabase()
	orderRepo := order_repo.NewOrderRepositorySQL(db)
	productRepo := product_repo.NewOrderRepositoryGRPC()
	orderService := order_serv.NewOrderService(orderRepo, productRepo, kafkaProducer)

	InitServiceGRPC := orderCtrl.NewOrderControllerGRPC(orderService)
	orderCtrl.NewOrderControllerKafka(orderService)

	s := grpc.NewServer()
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
	//kafka := pkg.NewKafkaProducer()
	//
	//orderRepo := order_repo.NewOrderRepositorySQL(db)
	//productRepo := product_repo.NewOrderRepositoryGRPC()
	//orderService := order_serv.NewOrderService(orderRepo, productRepo, kafka)
	//
	//orderCtrl.NewOrderControllerRestAPI(orderService, httpServer)
	//defaultCtrl.InitDefaultController(httpServer)
	//

	//log.Println("Starting REST server at", dto.CfgApp.RestPort)
	//err := httpServer.Run(fmt.Sprintf(`:%v`, dto.CfgApp.RestPort))
	//if err != nil {
	//	panic(err)
	//}

}
