package pkg

import (
	"fmt"
	ConsulAPI "github.com/hashicorp/consul/api"
	"service-product/dto"
	"service-product/utils"
)

func NewConsul(serviceName string, servicePort int, scheme string) {
	consulConf := ConsulAPI.DefaultConfig()
	consulConf.Address = fmt.Sprintf("%v:%v", dto.CfgConsul.ConsulHost, dto.CfgConsul.ConsulPort)
	consulConf.Scheme = "http"

	client, err := ConsulAPI.NewClient(consulConf)
	if err != nil {
		panic(err)
	}

	address := utils.GetLocalIP().String()
	registration := &ConsulAPI.AgentServiceRegistration{
		ID:      fmt.Sprintf("%v-%v-%v", serviceName, scheme, servicePort),
		Name:    serviceName,
		Port:    servicePort,
		Address: address,
		Check: &ConsulAPI.AgentServiceCheck{
			Interval:                       "10s",
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
