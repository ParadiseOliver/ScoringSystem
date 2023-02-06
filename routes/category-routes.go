package routes

import (
	"github.com/gin-gonic/gin"
)

type CategoryController interface {
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

func Categories(category *gin.RouterGroup, categoryController CategoryController) {

	category.GET("/disciplines", categoryController.AllDisciplines)
	category.POST("/discipline", categoryController.AddDiscipline)
	category.DELETE("/discipline/:id", categoryController.DelDiscipline)

	category.GET("/categories", categoryController.AllCategories)
	category.POST("/category", categoryController.AddCategory)
	category.DELETE("/category/:id", categoryController.DelCategory)

	category.GET("/agegroups", categoryController.AllAgeGroups)
	category.POST("/agegroup", categoryController.AddAgeGroup)
	category.DELETE("/agegroup/:id", categoryController.DelAgeGroup)

	category.GET("/genders", categoryController.AllGenders)
	category.POST("/gender", categoryController.AddGender)
	category.DELETE("/gender/:id", categoryController.DelGender)

	category.GET("/cat-groups", categoryController.AllCategoryGroups)
}
