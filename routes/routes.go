package routes

import (
	"customerlabs/business/events/repository"
	"customerlabs/business/events/usecase"
	eventsHTTP "customerlabs/business/events/webhttp"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine, eventChannel chan<- map[string]interface{}) {
	eventRepo := repository.NewEventRepo()
	eventUC := usecase.NewEventUC(eventRepo,eventChannel)
	eventsHTTP.NewHandler(router, eventUC)
}
