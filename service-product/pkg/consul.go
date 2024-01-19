package pkg

import (
	"fmt"
	ConsulAPI "github.com/hashicorp/consul/api"
	"service-product/dto"
)

func NewConsul(serviceName string, servicePort int) {
	consulConf := ConsulAPI.DefaultConfig()
	consulConf.Address = fmt.Sprintf("%v:%v", dto.CfgConsul.ConsulHost, dto.CfgConsul.ConsulPort)
	consulConf.Scheme = "http"

	client, err := ConsulAPI.NewClient(ConsulAPI.DefaultConfig())
	if err != nil {
		panic(err)
	}

	//address := utils.GetLocalIP().String()
	address := "localhost"
	fmt.Println("address ", address)

	registration := &ConsulAPI.AgentServiceRegistration{
		ID:      "IDNya",
		Name:    serviceName,
		Port:    dto.CfgApp.RestPort,
		Address: address,
		Check: &ConsulAPI.AgentServiceCheck{
			//GRPC:                           fmt.Sprintf("%s:%v/%s", "localhost", dto.CfgApp.GRPCPort, "grpc.health.v1.Health"),
			HTTP:                           fmt.Sprintf("http://%s:%v", address, dto.CfgApp.RestPort),
			Interval:                       "10s",
			Timeout:                        "10s",
			DeregisterCriticalServiceAfter: "5s",
		},
	}

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}

}
