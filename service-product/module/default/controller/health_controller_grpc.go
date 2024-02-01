package controller

import (
	"context"
	pb_user "service-product/proto/health"
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
