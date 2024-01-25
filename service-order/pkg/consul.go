package pkg

import (
	"fmt"
	ConsulAPI "github.com/hashicorp/consul/api"
	"service-order/dto"
	"service-order/utils"
)

func NewConsul(serviceName string, servicePort int, scheme string) {
	fmt.Println("NewConsul scheme", scheme)
	consulConf := ConsulAPI.DefaultConfig()
	consulConf.Address = fmt.Sprintf("%v:%v", dto.CfgConsul.ConsulHost, dto.CfgConsul.ConsulPort)
	consulConf.Scheme = "http"

	client, err := ConsulAPI.NewClient(consulConf)
	if err != nil {
		panic(err)
	}

	address := utils.GetLocalIP().String()
	registration := &ConsulAPI.AgentServiceRegistration{
		ID:      serviceName + scheme,
		Name:    serviceName + scheme,
		Port:    servicePort,
		Address: address,
		Check: &ConsulAPI.AgentServiceCheck{
			Interval:                       "30s",
			Timeout:                        "10s",
			DeregisterCriticalServiceAfter: "5s",
		},
	}
	if scheme == "GRPC" {
		registration.Check.GRPC = fmt.Sprintf("%s:%v/%v", address, dto.CfgApp.GRPCPort, "grpc.health.v1.Health")
	} else if scheme == "REST" {
		registration.Check.HTTP = fmt.Sprintf("http://%s:%v", address, dto.CfgApp.RestPort)
	}

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}

}
