package main

import (
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	hv1 "google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"net/http"
	pb "service-product/proto/product"
)

func main() {
	// Create a gRPC server
	grpcServer := grpc.NewServer()

	// Create a health check server
	healthServer := health.NewServer()
	healthServer.SetServingStatus("", hv1.HealthCheckResponse_SERVING)
	hv1.RegisterHealthServer(grpcServer, healthServer)

	RegisterGreeterService(grpcServer)

	// Start the gRPC server
	go func() {
		lis, err := net.Listen("tcp", ":8080")
		if err != nil {
			fmt.Println("Failed to listen:", err)
			return
		}
		fmt.Println("gRPC server listening on :8080")
		grpcServer.Serve(lis)
	}()

	// Register the service with Consul
	config := api.DefaultConfig()
	config.Address = "localhost:8500"
	config.Scheme = "http"
	client, err := api.NewClient(config)
	if err != nil {
		fmt.Println("Failed to create Consul client:", err)
		return
	}
	agent := client.Agent()

	registration := &api.AgentServiceRegistration{
		ID:      "id",
		Name:    "my-grpc-service",
		Address: "localhost",
		Port:    8081,
		Check: &api.AgentServiceCheck{
			HTTP: fmt.Sprintf("%s:%d", "localhost", 8081),
			//GRPC:                           fmt.Sprintf("%s:%d/%s", "localhost", 8080, "grpc.health.v1.Health"),
			Interval:                       "10s",
			DeregisterCriticalServiceAfter: "30s",
		},
	}

	err = agent.ServiceRegister(registration)
	if err != nil {
		fmt.Println("Failed to register service:", err)
		return
	}

	// Keep the process running
	http.ListenAndServe(":8081", nil)
}

type GreeterServer interface {
	Greet(context.Context, *pb.HelloRequest) (*pb.HelloReply, error)
}

// Implement the service
type greeterServer struct {
	pb.UnimplementedGreeterServer
}

func (s *greeterServer) Greet(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello, " + req.Name + "!"}, nil
}

// Register the service on the gRPC server
func RegisterGreeterService(grpcServer *grpc.Server) {
	pb.RegisterGreeterServer(grpcServer, &greeterServer{})
}
