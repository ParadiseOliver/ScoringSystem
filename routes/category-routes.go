package routes

import (
	"github.com/gin-gonic/gin"
)

type CategoryController interface {
	AllDisciplines(ctx *gin.Context)
	Discipline(ctx *gin.Context)
	AddDiscipline(ctx *gin.Context)
	UpdateDiscipline(ctx *gin.Context)
	DelDiscipline(ctx *gin.Context)
	AllCategories(ctx *gin.Context)
	Category(ctx *gin.Context)
	AddCategory(ctx *gin.Context)
	DelCategory(ctx *gin.Context)
	AllAgeGroups(ctx *gin.Context)
	AgeGroup(ctx *gin.Context)
	AddAgeGroup(ctx *gin.Context)
	DelAgeGroup(ctx *gin.Context)
	AllGenders(ctx *gin.Context)
	Gender(ctx *gin.Context)
	AddGender(ctx *gin.Context)
	DelGender(ctx *gin.Context)
	AllCategoryGroups(ctx *gin.Context)
	AddCategoryGroup(ctx *gin.Context)
	CategoryGroup(ctx *gin.Context)
	DelCategoryGroup(ctx *gin.Context)
}

func Categories(category *gin.RouterGroup, categoryController CategoryController) {

	category.GET("/disciplines", categoryController.AllDisciplines)
	category.GET("/discipline/:id", categoryController.Discipline)
	category.PUT("/discipline", categoryController.UpdateDiscipline)
	category.POST("/discipline", categoryController.AddDiscipline)
	category.DELETE("/discipline/:id", categoryController.DelDiscipline)

	category.GET("/categories", categoryController.AllCategories)
	category.GET("/category/:id", categoryController.Category)
	category.POST("/category", categoryController.AddCategory)
	category.DELETE("/category/:id", categoryController.DelCategory)

	category.GET("/agegroups", categoryController.AllAgeGroups)
	category.GET("/agegroup/:id", categoryController.AgeGroup)
	category.POST("/agegroup", categoryController.AddAgeGroup)
	category.DELETE("/agegroup/:id", categoryController.DelAgeGroup)

	category.GET("/genders", categoryController.AllGenders)
	category.GET("/gender/:id", categoryController.Gender)
	category.POST("/gender", categoryController.AddGender)
	category.DELETE("/gender/:id", categoryController.DelGender)

	category.GET("/cat-groups", categoryController.AllCategoryGroups)
	category.GET("/cat-groups/:id", categoryController.CategoryGroup)
	category.POST("/cat-groups", categoryController.AddCategoryGroup)
	category.DELETE("/cat-groups/:id", categoryController.DelCategoryGroup)
}
