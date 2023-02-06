package routes

import (
	"github.com/ParadiseOliver/ScoringSystem/controllers"
	"github.com/gin-gonic/gin"
)

func Categories(category *gin.RouterGroup, globalController controllers.CategoryController) {

	category.GET("/disciplines", globalController.AllDisciplines)
	category.POST("/discipline", globalController.AddDiscipline)
	category.DELETE("/discipline/:id", globalController.DelDiscipline)

	category.GET("/categories", globalController.AllCategories)
	category.POST("/category", globalController.AddCategory)
	category.DELETE("/category/:id", globalController.DelCategory)

	category.GET("/agegroups", globalController.AllAgeGroups)
	category.POST("/agegroup", globalController.AddAgeGroup)
	category.DELETE("/agegroup/:id", globalController.DelAgeGroup)

	category.GET("/genders", globalController.AllGenders)
	category.POST("/gender", globalController.AddGender)
	category.DELETE("/gender/:id", globalController.DelGender)

	category.GET("/cat-groups", globalController.AllCategoryGroups)
}
