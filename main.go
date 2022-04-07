package main

import (
	_ "api2022/lib/rpool"
	"api2022/server"
)

func main() {

	server.RunHttpServer()

}
