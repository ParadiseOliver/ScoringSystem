package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/ParadiseOliver/ScoringSystem/entity"
	"github.com/gin-gonic/gin"
)

type ResultsService interface {
	AllResultsByEventId(id string) ([]entity.Result, error)
	ResultByResultId(id string) (*entity.Result, error)
	ResultsByAthleteId(id string) ([]entity.Result, error)
	UserByUserId(id string) (*entity.User, error)
	ScoreAthlete(eventId, athleteId int, score *entity.TriScore) (*entity.Result, error)
}

type resultsController struct {
	service ResultsService
}

func NewResultsController(service ResultsService) *resultsController {
	//validate := validator.New()
	//validate.RegisterValidation("is-after", validators.ValidateIsAfter)
	return &resultsController{
		service: service,
	}
}

func (c *resultsController) AllResultsByEventId(ctx *gin.Context) {
	id := ctx.Param("eventId")
	results, err := c.service.AllResultsByEventId(id)

	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}
	wrapped := struct {
		Fields []entity.Result `json:"results"`
	}{
		results,
	}
	ctx.JSON(http.StatusOK, wrapped)
}

func (c *resultsController) ResultByResultId(ctx *gin.Context) {
	id := ctx.Param("resultId")
	result, err := c.service.ResultByResultId(id)

	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}
	wrapped := struct {
		Fields entity.Result `json:"result"`
	}{
		*result,
	}
	ctx.JSON(http.StatusOK, wrapped)
}

func (c *resultsController) ResultsByAthleteId(ctx *gin.Context) {
	athleteId := ctx.Param("athleteId")
	results, err := c.service.ResultsByAthleteId(athleteId)

	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}
	wrapped := struct {
		Fields []entity.Result `json:"results"`
	}{
		results,
	}
	ctx.JSON(http.StatusOK, wrapped)
}

func (c *resultsController) UserByUserId(ctx *gin.Context) {
	id := ctx.Param("athleteId")
	user, err := c.service.UserByUserId(id)

	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}
	data := gin.H{
		"user": user,
	}
	ctx.HTML(http.StatusOK, "score.html", data)
}

func (c *resultsController) ScoreAthlete(ctx *gin.Context) {
	eventId, err := strconv.Atoi(ctx.Param("eventId"))
	if err != nil {
		log.Printf("Failed to get eventID: %v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	athleteId, err := strconv.Atoi(ctx.Param("athleteId"))
	if err != nil {
		log.Printf("Failed to add athleteID: %v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	var score *entity.TriScore
	err = ctx.ShouldBind(&score)
	if err != nil {
		log.Printf("Failed to bind score: %v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	score.Total = score.E1 + score.E2 + score.HD + score.DD + score.Tof - score.Pen // Just summing first 2 E scores rather than median currently.

	var result *entity.Result
	result, err = c.service.ScoreAthlete(eventId, athleteId, score)

	if err != nil {
		log.Printf("Failed to add score: %v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	wrapped := struct {
		Fields entity.Result `json:"score"`
	}{
		*result,
	}
	ctx.JSON(http.StatusOK, wrapped)
}
