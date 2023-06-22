package controllers

import (
	"log"
	"net/http"

	"github.com/ParadiseOliver/ScoringSystem/entity"
	"github.com/gin-gonic/gin"
)

type CategoryService interface {
	AllDisciplines() ([]entity.Discipline, error)
	Discipline(id string) (*entity.Discipline, error)
	AddDiscipline(discipline *entity.Discipline) (*entity.Discipline, error)
	UpdateDiscipline(discipline *entity.Discipline) (*entity.Discipline, error)
	DelDiscipline(id string) error
	AllCategories() ([]entity.Category, error)
	Category(id string) (*entity.Category, error)
	AddCategory(category *entity.Category) (*entity.Category, error)
	DelCategory(id string) error
	AllAgeGroups() ([]entity.AgeGroup, error)
	AgeGroup(id string) (*entity.AgeGroup, error)
	AddAgeGroup(ageGroup *entity.AgeGroup) (*entity.AgeGroup, error)
	DelAgeGroup(id string) error
	AllGenders() ([]entity.Gender, error)
	Gender(id string) (*entity.Gender, error)
	AddGender(gender *entity.Gender) (*entity.Gender, error)
	DelGender(id string) error
	AllCategoryGroups() ([]entity.CategoryGroup, error)
	AddCategoryGroup(catGroup *entity.CategoryGroup) (*entity.CategoryGroup, error)
	CategoryGroup(id string) (*entity.CategoryGroup, error)
	DelCategoryGroup(id string) error
}

type categoryController struct {
	service CategoryService
}

func NewCategoryController(service CategoryService) *categoryController {
	//validate := validator.New()
	//validate.RegisterValidation("is-after", validators.ValidateIsAfter)
	return &categoryController{
		service: service,
	}
}

func (c *categoryController) AllDisciplines(ctx *gin.Context) {
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

func (c *categoryController) Discipline(ctx *gin.Context) {
	id := ctx.Param("id")
	discipline, err := c.service.Discipline(id)

	if err != nil {
		log.Printf("Failed to get discipline: %v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	wrapped := struct {
		Fields entity.Discipline `json:"discipline"`
	}{
		*discipline,
	}
	ctx.JSON(http.StatusOK, wrapped)
}

func (c *categoryController) AddDiscipline(ctx *gin.Context) {
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

func (c *categoryController) UpdateDiscipline(ctx *gin.Context) {
	var discipline *entity.Discipline
	err := ctx.ShouldBindJSON(&discipline)
	if err != nil {
		log.Printf("Failed to bind discipline: %v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	discipline, err = c.service.UpdateDiscipline(discipline)
	if err != nil {
		log.Printf("Failed to update discipline: %v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, discipline)
}

func (c *categoryController) DelDiscipline(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.service.DelDiscipline(id)
	if err != nil {
		log.Printf("Failed to delete discipline: %v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.Status(http.StatusOK)
}

func (c *categoryController) AllCategories(ctx *gin.Context) {
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

func (c *categoryController) Category(ctx *gin.Context) {
	id := ctx.Param("id")
	cat, err := c.service.Category(id)

	if err != nil {
		log.Printf("Failed to get category: %v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	wrapped := struct {
		Fields entity.Category `json:"category"`
	}{
		*cat,
	}
	ctx.JSON(http.StatusOK, wrapped)
}

func (c *categoryController) AddCategory(ctx *gin.Context) {
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

func (c *categoryController) DelCategory(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.service.DelCategory(id)
	if err != nil {
		log.Printf("Failed to delete category: %v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.Status(http.StatusOK)
}

func (c *categoryController) AllAgeGroups(ctx *gin.Context) {
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

func (c *categoryController) AgeGroup(ctx *gin.Context) {
	id := ctx.Param("id")
	age, err := c.service.AgeGroup(id)

	if err != nil {
		log.Printf("Failed to get agegroup: %v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	wrapped := struct {
		Fields entity.AgeGroup `json:"agegroup"`
	}{
		*age,
	}
	ctx.JSON(http.StatusOK, wrapped)
}

func (c *categoryController) AddAgeGroup(ctx *gin.Context) {
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

func (c *categoryController) DelAgeGroup(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.service.DelAgeGroup(id)
	if err != nil {
		log.Printf("Failed to delete age group: %v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.Status(http.StatusOK)
}

func (c *categoryController) AllGenders(ctx *gin.Context) {
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

func (c *categoryController) Gender(ctx *gin.Context) {
	id := ctx.Param("id")
	gender, err := c.service.Gender(id)

	if err != nil {
		log.Printf("Failed to get gender: %v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	wrapped := struct {
		Fields entity.Gender `json:"gender"`
	}{
		*gender,
	}
	ctx.JSON(http.StatusOK, wrapped)
}

func (c *categoryController) AddGender(ctx *gin.Context) {
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

func (c *categoryController) DelGender(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.service.DelGender(id)
	if err != nil {
		log.Printf("Failed to delete gender: %v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.Status(http.StatusOK)
}

func (c *categoryController) AllCategoryGroups(ctx *gin.Context) {
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

func (c *categoryController) CategoryGroup(ctx *gin.Context) {
	id := ctx.Param("id")
	categoryGroup, err := c.service.CategoryGroup(id)

	if err != nil {
		log.Printf("Failed to get category group: %v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	wrapped := struct {
		Fields entity.CategoryGroup `json:"category_group"`
	}{
		*categoryGroup,
	}
	ctx.JSON(http.StatusOK, wrapped)
}

func (c *categoryController) DelCategoryGroup(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.service.DelCategoryGroup(id)
	if err != nil {
		log.Printf("Failed to delete category group: %v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.Status(http.StatusOK)
}

func (c *categoryController) AddCategoryGroup(ctx *gin.Context) {
	var catGroup *entity.CategoryGroup
	err := ctx.ShouldBindJSON(&catGroup)
	if err != nil {
		log.Printf("Failed to bind gender: %v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	catGroup, err = c.service.AddCategoryGroup(catGroup)
	if err != nil {
		log.Printf("Failed to add gender: %v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, catGroup)
}
