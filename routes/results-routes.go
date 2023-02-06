package routes

import "github.com/gin-gonic/gin"

type ResultsController interface {
	ResultByResultId(ctx *gin.Context)
	AllResultsByEventId(ctx *gin.Context)
	ResultsByAthleteId(ctx *gin.Context)
}

func Results(results *gin.RouterGroup, resultController ResultsController) {
	results.GET("/:resultId", resultController.ResultByResultId)
	results.GET("/event/:eventId", resultController.AllResultsByEventId)
	results.GET("/athlete/:athleteId", resultController.ResultsByAthleteId)
}
