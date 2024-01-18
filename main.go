package main

import (
	"customerlabs/business/events/usecase"
	"customerlabs/routes"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("CUSTOMERLABS ASSIGNMENT-VEDANT")
	requestChannel := make(chan map[string]interface{}, 100)
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	go usecase.Worker(requestChannel, shutdown)

	router := gin.Default()
	routes.InitRoutes(router, requestChannel)
	go func() {
		err := router.Run(":8080")
		if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	}()

	<-shutdown
	close(requestChannel)
}