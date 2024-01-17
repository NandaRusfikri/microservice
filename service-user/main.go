package main

import "service-user/app"

func main() {

	app.NewGRPC()
	app.NewRestAPI()

}
