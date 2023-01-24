package controllers

import (
	"net/http"

	"github.com/ParadiseOliver/ScoringSystem/entity"
	"github.com/ParadiseOliver/ScoringSystem/usecases"
	"github.com/ParadiseOliver/ScoringSystem/validators"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type EventController interface {
	GetAll() ([]entity.Event, error)
	Create(ctx *gin.Context) (entity.Event, error)
	AllEvents(ctx *gin.Context)
}

type controller struct {
	service usecases.EventService
}

var (
	validate *validator.Validate
)

func New(service usecases.EventService) EventController {
	validate = validator.New()
	validate.RegisterValidation("is-after", validators.ValidateIsAfter)
	return &controller{
		service: service,
	}
}

func (c *controller) Create(ctx *gin.Context) (entity.Event, error) {
	var event entity.Event
	err := ctx.ShouldBindJSON(&event)
	if err != nil {
		return entity.Event{}, err
	}
	err = validate.Struct(event)
	if err != nil {
		return entity.Event{}, err
	}
	c.service.Create(event)
	return event, nil
}

func (c *controller) GetAll() ([]entity.Event, error) {
	return c.service.GetAll()
}

func (c *controller) AllEvents(ctx *gin.Context) {
	events, err := c.service.GetAll()
	if err != nil {
		panic(err)
	}
	data := gin.H{
		"title":  "Scoring System",
		"events": events,
	}
	ctx.HTML(http.StatusOK, "all_events.html", data)
}
