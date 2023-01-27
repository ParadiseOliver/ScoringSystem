package controllers

import (
	"log"
	"net/http"

	"github.com/ParadiseOliver/ScoringSystem/entity"
	"github.com/ParadiseOliver/ScoringSystem/usecases"
	"github.com/ParadiseOliver/ScoringSystem/validators"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type EventController interface {
	GetAll(ctx *gin.Context)
	CreateEvent(ctx *gin.Context)
	AllEvents(ctx *gin.Context)
	GetEventById(ctx *gin.Context)
	AllResultsByEventId(ctx *gin.Context)
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

func (c *controller) CreateEvent(ctx *gin.Context) {
	var event *entity.Event
	err := ctx.ShouldBindJSON(&event)
	if err != nil {
		log.Fatal(err)
	}
	err = validate.Struct(event)
	if err != nil {
		log.Fatal(err)
	}
	event, err = c.service.CreateEvent(event)
	if err != nil {
		log.Fatal(err)
	}
	ctx.IndentedJSON(http.StatusOK, event)
}

func (c *controller) GetAll(ctx *gin.Context) {
	events, err := c.service.GetAll()
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx.IndentedJSON(http.StatusOK, events)
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

func (c *controller) GetEventById(ctx *gin.Context) {
	id := ctx.Param("eventId")
	event, err := c.service.GetEventById(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Event not found"})
	}
	ctx.IndentedJSON(http.StatusOK, event)
}

func (c *controller) AllResultsByEventId(ctx *gin.Context) {
	id := ctx.Param("eventId")
	results, err := c.service.AllResultsByEventId(id)

	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Results not found"})
		return
	}
	ctx.IndentedJSON(http.StatusOK, results)
}
