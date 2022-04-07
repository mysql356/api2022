package server

import (
	"api2022/routers"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func RunHttpServer() {

	routers.SetupRouter()
	log.Printf("System Ready at %s \n", viper.GetString("Port"))
	http.ListenAndServe(":"+viper.GetString("Port"), nil)

}
