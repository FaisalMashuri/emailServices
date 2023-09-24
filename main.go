package main

import (
	"fmt"
	"github.com/FaisalMashuri/emailServices/config/env"
	"github.com/FaisalMashuri/emailServices/config/rabbitmq"
	"github.com/FaisalMashuri/emailServices/routes"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := env.Config
	h := gin.Default()
	rabbbitmq := rabbitmq.NewRabbitMQ(cfg)
	routes.NewRoutes(cfg, h, rabbbitmq)

	err := h.Run(fmt.Sprintf(":%s", cfg.Port))
	if err != nil {
		log.Fatal(err.Error())
	}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down...")

}
