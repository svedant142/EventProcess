package webhttp

import (
	"customerlabs/business/events/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)
type Handler struct {
	uc usecase.IEventUC
}

func NewHandler(router *gin.Engine, eventuc usecase.IEventUC) {
	handler := Handler{uc: eventuc}
	Init(router, handler)
}

func Init(router *gin.Engine, h Handler) {
	router.POST("/event-process", h.ProcessEvent)
}


func (h *Handler) ProcessEvent(c *gin.Context) {
	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req) ; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : "invalid payload",
			"status" : false,
		})
		return
	}
	h.uc.ProcessEvent(c, req)

	c.JSON(http.StatusOK, gin.H{"message": "request received and being processed"})
}
