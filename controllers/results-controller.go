package controllers

import (
	"net/http"

	"github.com/ParadiseOliver/ScoringSystem/entity"
	"github.com/gin-gonic/gin"
)

type ResultsService interface {
	AllResultsByEventId(id string) ([]entity.Result, error)
	ResultByResultId(id string) (*entity.Result, error)
	ResultsByAthleteId(id string) ([]entity.Result, error)
	UserByUserId(id string) (*entity.User, error)
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
