package delivery

import (
	"github.com/ParadiseOliver/ScoringSystem/entity"
	"github.com/ParadiseOliver/ScoringSystem/usecases"
	"github.com/gin-gonic/gin"
)

type EventController interface {
	GetAll() []entity.Event
	Create(ctx *gin.Context) entity.Event
}

type controller struct {
	service usecases.EventService
}

func New(service usecases.EventService) EventController {
	return &controller{
		service: service,
	}
}

func (c *controller) Create(ctx *gin.Context) entity.Event {
	var event entity.Event
	ctx.ShouldBindJSON(&event)
	c.service.Create(event)
	return event
}

func (c *controller) GetAll() []entity.Event {
	return c.service.GetAll()
}
