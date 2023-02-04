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

type EventService interface { // TODO: Move me to transport layer (currently main)
	GetAll() ([]entity.Event, error)
	CreateEvent(event *entity.Event) (*entity.Event, error)
	GetEventById(id string) (*entity.Event, error)
	AllResultsByEventId(id string) ([]entity.Result, error)
	ResultByResultId(id string) (*entity.Result, error)
	ResultsByAthleteId(id string) ([]entity.Result, error)
}

type EventController interface {
	GetAll(ctx *gin.Context)
	CreateEvent(ctx *gin.Context)
	AllEvents(ctx *gin.Context)
	GetEventById(ctx *gin.Context)
	AllResultsByEventId(ctx *gin.Context)
	ResultByResultId(ctx *gin.Context)
	ResultsByAthleteId(ctx *gin.Context)
}

type eventController struct {
	service EventService
}

func New(service EventService) EventController {
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
	ctx.JSON(http.StatusOK, events)
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

func (c *eventController) AllEvents(ctx *gin.Context) {
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
	ctx.HTML(http.StatusOK, "all_events.html", data)
}

func (c *eventController) GetEventById(ctx *gin.Context) {
	id := ctx.Param("eventId")
	event, err := c.service.GetEventById(id)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
	}
	ctx.JSON(http.StatusOK, event)
}

func (c *eventController) AllResultsByEventId(ctx *gin.Context) {
	id := ctx.Param("eventId")
	results, err := c.service.AllResultsByEventId(id)

	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, results)
}

func (c *eventController) ResultByResultId(ctx *gin.Context) {
	id := ctx.Param("resultId")
	result, err := c.service.ResultByResultId(id)

	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (c *eventController) ResultsByAthleteId(ctx *gin.Context) {
	athleteId := ctx.Param("athleteId")
	results, err := c.service.ResultsByAthleteId(athleteId)

	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, results)
}
