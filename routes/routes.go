package routes

import (
	"fmt"

	"github.com/FaisalMashuri/emailServices/Controllers"
	"github.com/FaisalMashuri/emailServices/config/rabbitmq"
	"github.com/FaisalMashuri/emailServices/domain/usecase"
	"github.com/FaisalMashuri/emailServices/models"
	"github.com/gin-gonic/gin"
)

type RouteParams struct {
	config models.AppConfig
	c      *gin.Engine
}

func NewRoutes(cfg models.AppConfig, c *gin.Engine, consumer *rabbitmq.RabbitMQ) {
	v1 := c.Group("api/v1")
	v1.Static("/static", "/static")

	fmt.Println(cfg.EmailConfig)
	emailServices := usecase.NewEmailService(cfg.EmailConfig)
	Controllers.NewEmailController(v1, emailServices)
	go consumer.EmailConsumer(emailServices)

}
