package controllers

import (
	"log"
	"net/http"

	"github.com/ParadiseOliver/ScoringSystem/entity"
	"github.com/gin-gonic/gin"
)

type GlobalController interface {
	AllDisciplines(ctx *gin.Context)
	AddDiscipline(ctx *gin.Context)
	DelDiscipline(ctx *gin.Context)
	AllCategories(ctx *gin.Context)
	AllAgeGroups(ctx *gin.Context)
	AllGenders(ctx *gin.Context)
	AllCategoryGroups(ctx *gin.Context)
}

type GlobalService interface {
	AllDisciplines() ([]entity.Discipline, error)
	AddDiscipline(discipline *entity.Discipline) (*entity.Discipline, error)
	DelDiscipline(id string) error
	AllCategories() ([]entity.Category, error)
	AllAgeGroups() ([]entity.AgeGroup, error)
	AllGenders() ([]entity.Gender, error)
	AllCategoryGroups() ([]entity.CategoryGroup, error)
}

type globalController struct {
	service GlobalService
}

func NewGlobalController(service GlobalService) GlobalController {
	//validate := validator.New()
	//validate.RegisterValidation("is-after", validators.ValidateIsAfter)
	return &globalController{
		service: service,
	}
}

func (c *globalController) AllDisciplines(ctx *gin.Context) {
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

func (c *globalController) AddDiscipline(ctx *gin.Context) {
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

func (c *globalController) DelDiscipline(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.service.DelDiscipline(id)
	if err != nil {
		log.Printf("Failed to delete discipline: %v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.Status(http.StatusOK)
}

func (c *globalController) AllCategories(ctx *gin.Context) {
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

func (c *globalController) AllAgeGroups(ctx *gin.Context) {
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

func (c *globalController) AllGenders(ctx *gin.Context) {
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

func (c *globalController) AllCategoryGroups(ctx *gin.Context) {
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
