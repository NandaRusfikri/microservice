package main

import (
	"service-order/app"
)

func main() {

	go app.NewGRPC()
	app.NewRestAPI()

}
