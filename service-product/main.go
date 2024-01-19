package main

import (
	"service-product/app"
)

func main() {

	go app.NewGRPC()
	app.NewRestAPI()

}
