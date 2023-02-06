package routes

import (
	"github.com/ParadiseOliver/ScoringSystem/controllers"
	"github.com/gin-gonic/gin"
)

func Categories(events *gin.RouterGroup, globalController controllers.GlobalController) {

	events.GET("/disciplines", globalController.AllDisciplines)
	events.POST("/discipline", globalController.AddDiscipline)
	events.DELETE("/discipline/:id", globalController.DelDiscipline)

	events.GET("/categories", globalController.AllCategories)
	events.POST("/category", globalController.AddCategory)
	events.DELETE("/category/:id", globalController.DelCategory)

	events.GET("/agegroups", globalController.AllAgeGroups)
	events.POST("/agegroup", globalController.AddAgeGroup)
	events.DELETE("/agegroup/:id", globalController.DelAgeGroup)

	events.GET("/genders", globalController.AllGenders)
	events.POST("/gender", globalController.AddGender)
	events.DELETE("/gender/:id", globalController.DelGender)

	events.GET("/cat-groups", globalController.AllCategoryGroups)
}
