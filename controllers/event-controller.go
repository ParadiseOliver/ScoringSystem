package controllers

import (
	"log"
	"net/http"

	"github.com/ParadiseOliver/ScoringSystem/entity"
	"github.com/ParadiseOliver/ScoringSystem/usecases"
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

type EventController interface {
	GetAll(ctx *gin.Context)
	CreateEvent(ctx *gin.Context)
	AllEvents(ctx *gin.Context)
	GetEventById(ctx *gin.Context)
	AllResultsByEventId(ctx *gin.Context)
	ResultByResultId(ctx *gin.Context)
	ResultsByAthleteId(ctx *gin.Context)
	AllDisciplines(ctx *gin.Context)
	AddDiscipline(ctx *gin.Context)
	AllCategories(ctx *gin.Context)
	AllAgeGroups(ctx *gin.Context)
	AllGenders(ctx *gin.Context)
	AllCategoryGroups(ctx *gin.Context)
}

type controller struct {
	service usecases.EventService
}

func New(service usecases.EventService) EventController {
	//validate := validator.New()
	//validate.RegisterValidation("is-after", validators.ValidateIsAfter)
	return &controller{
		service: service,
	}
}

func (c *controller) GetAll(ctx *gin.Context) {
	events, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, events)
}

func (c *controller) CreateEvent(ctx *gin.Context) {
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

func (c *controller) AllEvents(ctx *gin.Context) {
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

func (c *controller) GetEventById(ctx *gin.Context) {
	id := ctx.Param("eventId")
	event, err := c.service.GetEventById(id)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
	}
	ctx.JSON(http.StatusOK, event)
}

func (c *controller) AllResultsByEventId(ctx *gin.Context) {
	id := ctx.Param("eventId")
	results, err := c.service.AllResultsByEventId(id)

	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, results)
}

func (c *controller) ResultByResultId(ctx *gin.Context) {
	id := ctx.Param("resultId")
	result, err := c.service.ResultByResultId(id)

	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (c *controller) ResultsByAthleteId(ctx *gin.Context) {
	athleteId := ctx.Param("athleteId")
	results, err := c.service.ResultsByAthleteId(athleteId)

	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, results)
}

func (c *controller) AllDisciplines(ctx *gin.Context) {
	disciplines, err := c.service.AllDisciplines()

	if err != nil {
		log.Printf("Failed to get disciplines: %v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	wrapped := struct {
		Fields []entity.Discipline `json:"disciplines"`
	}{
		disciplines,
	}
	ctx.JSON(http.StatusOK, wrapped)
}

func (c *controller) AddDiscipline(ctx *gin.Context) {
	var discipline *entity.Discipline
	err := ctx.ShouldBindJSON(&discipline)
	if err != nil {
		log.Printf("Failed to add discipline: %v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	discipline, err = c.service.AddDiscipline(discipline)
	if err != nil {
		log.Printf("Failed to add discipline: %v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, discipline)
}

func (c *controller) AllCategories(ctx *gin.Context) {
	categories, err := c.service.AllCategories()

	if err != nil {
		log.Printf("Failed to get categories: %v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	wrapped := struct {
		Fields []entity.Category `json:"categories"`
	}{
		categories,
	}
	ctx.JSON(http.StatusOK, wrapped)
}

func (c *controller) AllAgeGroups(ctx *gin.Context) {
	ageGroups, err := c.service.AllAgeGroups()

	if err != nil {
		log.Printf("Failed to get age groups: %v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	wrapped := struct {
		Fields []entity.AgeGroup `json:"agegroups"`
	}{
		ageGroups,
	}
	ctx.JSON(http.StatusOK, wrapped)
}

func (c *controller) AllGenders(ctx *gin.Context) {
	genders, err := c.service.AllGenders()

	if err != nil {
		log.Printf("Failed to get genders: %v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	wrapped := struct {
		Fields []entity.Gender `json:"genders"`
	}{
		genders,
	}
	ctx.JSON(http.StatusOK, wrapped)
}

func (c *controller) AllCategoryGroups(ctx *gin.Context) {
	categoryGroups, err := c.service.AllCategoryGroups()

	if err != nil {
		log.Printf("Failed to get category groups: %v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	wrapped := struct {
		Fields []entity.CategoryGroup `json:"category_groups"`
	}{
		categoryGroups,
	}
	ctx.JSON(http.StatusOK, wrapped)
}
