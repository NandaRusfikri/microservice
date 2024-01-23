package util

import "go-micro.dev/v4/client"

var clients = client.DefaultClient
//var address := []string{"localhost:34567"}
//var opts = client.Options{CallOptions: client.CallOptions{Address: address}}



func InitClient(input client.Client){
	clients = input



}

func GetClient()client.Client{
	return clients
}
