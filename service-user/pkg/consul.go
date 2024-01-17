package pkg

import (
	"fmt"
	ConsulAPI "github.com/hashicorp/consul/api"
	"service-user/dto"
)

func NewConsul(config dto.ConfigConsul) {
	consulConf := ConsulAPI.DefaultConfig()
	consulConf.Address = fmt.Sprintf("%v:%v", config.ConsulHost, config.ConsulPort)
	consulConf.Scheme = "http"

	client, err := ConsulAPI.NewClient(ConsulAPI.DefaultConfig())
	if err != nil {
		panic(err)
	}

	registration := &ConsulAPI.AgentServiceRegistration{
		Name: config.ServiceName,
		Port: config.ConsulPort,
	}

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}
}
