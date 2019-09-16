package main

import (
	"app/config"
	"app/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(config.ServerConfig.RunMode)

	routersInit := routes.InitRouter()
	readTimeout := config.ServerConfig.ReadTimeout
	writeTimeout := config.ServerConfig.WriteTimeout
	endPoint := fmt.Sprintf(":%d", config.ServerConfig.HttpPort)

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}






