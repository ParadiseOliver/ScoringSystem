package routes

import "github.com/gin-gonic/gin"

type EventController interface {
	GetAll(ctx *gin.Context)
	GetEventById(ctx *gin.Context)
	CreateEvent(ctx *gin.Context)
}

func Events(events *gin.RouterGroup, eventController EventController) {
	events.GET("/", eventController.GetAll)
	events.GET("/:eventId", eventController.GetEventById)
	events.POST("/", eventController.CreateEvent)
}
