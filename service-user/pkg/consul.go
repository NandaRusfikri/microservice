package pkg

import (
	"fmt"
	ConsulAPI "github.com/hashicorp/consul/api"
	"service-user/dto"
)

func NewConsul(serviceName string, servicePort int) {
	consulConf := ConsulAPI.DefaultConfig()
	consulConf.Address = fmt.Sprintf("%v:%v", dto.CfgConsul.ConsulHost, dto.CfgConsul.ConsulPort)
	consulConf.Scheme = "http"

	client, err := ConsulAPI.NewClient(ConsulAPI.DefaultConfig())
	if err != nil {
		panic(err)
	}

	registration := &ConsulAPI.AgentServiceRegistration{
		Name: serviceName,
		Port: servicePort,
	}

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}
}
