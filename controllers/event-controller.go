package controllers

import (
	"log"
	"net/http"

	"github.com/ParadiseOliver/ScoringSystem/entity"
	"github.com/gin-gonic/gin"
	logger "go.uber.org/zap"
)

func init() {
	logger, err := logger.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()
	//sugar := logger.Sugar()
}

type EventService interface {
	GetAll() ([]entity.Event, error)
	CreateEvent(event *entity.Event) (*entity.Event, error)
	GetEventById(id string) (*entity.Event, error)
}

type eventController struct {
	service EventService
}

func NewEventController(service EventService) *eventController {
	//validate := validator.New()
	//validate.RegisterValidation("is-after", validators.ValidateIsAfter)
	return &eventController{
		service: service,
	}
}

func (c *eventController) GetAll(ctx *gin.Context) {
	events, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	wrapped := struct {
		Fields []entity.Event `json:"events"`
	}{
		events,
	}
	ctx.JSON(http.StatusOK, wrapped)
}

func (c *eventController) CreateEvent(ctx *gin.Context) {
	var event *entity.Event
	err := ctx.ShouldBindJSON(&event)
	if err != nil {
		log.Printf("Failed to get all events: %v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	//err = validate.Struct(event)
	//if err != nil {
	//	log.Fatal(err)
	//}
	event, err = c.service.CreateEvent(event)
	if err != nil {
		log.Printf("Failed to get all events: %v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, event)
}

func (c *eventController) EventsPage(ctx *gin.Context) {
	events, err := c.service.GetAll()
	if err != nil {
		log.Printf("Failed to get all events: %v", err) // This would be log.Errorf for example
		ctx.Status(http.StatusInternalServerError)      // Status 500 because we don't want to expose our internal workings to a bad actor.
		return                                          // Return so we don't try and continue with the rest of the logic..
	}
	data := gin.H{
		"title":  "Scoring System",
		"events": events,
	}
	ctx.HTML(http.StatusOK, "allEvents.html", data)
}

func (c *eventController) GetEventById(ctx *gin.Context) {
	id := ctx.Param("eventId")
	event, err := c.service.GetEventById(id)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
	}
	ctx.JSON(http.StatusOK, event)
}

func (c *eventController) EventPage(ctx *gin.Context) {
	id := ctx.Param("eventId")
	event, err := c.service.GetEventById(id)
	if err != nil {
		log.Printf("Failed to get event: %v", err) // This would be log.Errorf for example
		ctx.Status(http.StatusInternalServerError) // Status 500 because we don't want to expose our internal workings to a bad actor.
		return                                     // Return so we don't try and continue with the rest of the logic..
	}
	data := gin.H{
		"title": "Scoring System",
		"event": event,
	}
	ctx.HTML(http.StatusOK, "event.html", data)
}
