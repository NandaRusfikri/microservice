package main

import "service-user/app"

func main() {

	go app.NewGRPC()
	app.NewRestAPI()

}
