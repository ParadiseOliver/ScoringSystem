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
	AddCategory(ctx *gin.Context)
	DelCategory(ctx *gin.Context)
	AllAgeGroups(ctx *gin.Context)
	AddAgeGroup(ctx *gin.Context)
	DelAgeGroup(ctx *gin.Context)
	AllGenders(ctx *gin.Context)
	AddGender(ctx *gin.Context)
	DelGender(ctx *gin.Context)
	AllCategoryGroups(ctx *gin.Context)
}

type GlobalService interface {
	AllDisciplines() ([]entity.Discipline, error)
	AddDiscipline(discipline *entity.Discipline) (*entity.Discipline, error)
	DelDiscipline(id string) error
	AllCategories() ([]entity.Category, error)
	AddCategory(category *entity.Category) (*entity.Category, error)
	DelCategory(id string) error
	AllAgeGroups() ([]entity.AgeGroup, error)
	AddAgeGroup(ageGroup *entity.AgeGroup) (*entity.AgeGroup, error)
	DelAgeGroup(id string) error
	AllGenders() ([]entity.Gender, error)
	AddGender(gender *entity.Gender) (*entity.Gender, error)
	DelGender(id string) error
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
		log.Printf("Failed to bind discipline: %v", err)
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

func (c *globalController) AddCategory(ctx *gin.Context) {
	var category *entity.Category
	err := ctx.ShouldBindJSON(&category)
	if err != nil {
		log.Printf("Failed to bind category: %v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	category, err = c.service.AddCategory(category)
	if err != nil {
		log.Printf("Failed to add category: %v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, category)
}

func (c *globalController) DelCategory(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.service.DelCategory(id)
	if err != nil {
		log.Printf("Failed to delete category: %v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.Status(http.StatusOK)
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

func (c *globalController) AddAgeGroup(ctx *gin.Context) {
	var ageGroup *entity.AgeGroup
	err := ctx.ShouldBindJSON(&ageGroup)
	if err != nil {
		log.Printf("Failed to bind age group: %v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ageGroup, err = c.service.AddAgeGroup(ageGroup)
	if err != nil {
		log.Printf("Failed to add age group: %v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, ageGroup)
}

func (c *globalController) DelAgeGroup(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.service.DelAgeGroup(id)
	if err != nil {
		log.Printf("Failed to delete age group: %v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.Status(http.StatusOK)
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

func (c *globalController) AddGender(ctx *gin.Context) {
	var gender *entity.Gender
	err := ctx.ShouldBindJSON(&gender)
	if err != nil {
		log.Printf("Failed to bind gender: %v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	gender, err = c.service.AddGender(gender)
	if err != nil {
		log.Printf("Failed to add gender: %v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, gender)
}

func (c *globalController) DelGender(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.service.DelGender(id)
	if err != nil {
		log.Printf("Failed to delete gender: %v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.Status(http.StatusOK)
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
