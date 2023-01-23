package delivery

import (
	"net/http"

	"github.com/ParadiseOliver/ScoringSystem/entity"
	"github.com/ParadiseOliver/ScoringSystem/usecases"
	"github.com/ParadiseOliver/ScoringSystem/validators"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type EventController interface {
	GetAll() []entity.Event
	Create(ctx *gin.Context) error
	AllEvents(ctx *gin.Context)
}

type controller struct {
	service usecases.EventService
}

var validate *validator.Validate

func New(service usecases.EventService) EventController {
	validate = validator.New()
	validate.RegisterValidation("is-after", validators.ValidateIsAfter)
	return &controller{
		service: service,
	}
}

func (c *controller) Create(ctx *gin.Context) error {
	var event entity.Event
	err := ctx.ShouldBindJSON(&event)
	if err != nil {
		return err
	}
	err = validate.Struct(event)
	if err != nil {
		return err
	}
	c.service.Create(event)
	return nil
}

func (c *controller) GetAll() []entity.Event {
	return c.service.GetAll()
}

func (c *controller) AllEvents(ctx *gin.Context) {
	events := c.service.GetAll()
	data := gin.H{
		"title":  "Scoring System",
		"events": events,
	}
	ctx.HTML(http.StatusOK, "all_events.html", data)
}