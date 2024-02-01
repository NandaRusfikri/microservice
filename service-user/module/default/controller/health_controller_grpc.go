package controller

import (
	"context"
	pb_user "service-user/proto/user"
)

type healthCheck struct {
	pb_user.UnimplementedHealthServer
}

func NewhealthCheck() *healthCheck {
	return &healthCheck{}
}

func (h *healthCheck) Check(ctx context.Context, req *pb_user.HealthCheckRequest) (*pb_user.HealthCheckResponse, error) {
	var res pb_user.HealthCheckResponse
	res.Status = pb_user.HealthCheckResponse_SERVING

	return &res, nil
}

//func (h *healthCheck) Watch(req *pb_user.HealthCheckRequest, stream pb_user.Health_WatchServer) error {
//
//	//for {
//	//	time.Sleep(5 * time.Second)
//	//	err := stream.Send(&pb_user.HealthCheckResponse{Status: pb_user.HealthCheckResponse_SERVING})
//	//	if err != nil {
//	//		return err
//	//	}
//	//	time.Sleep(5 * time.Second)
//	//}
//	return nil
//}
