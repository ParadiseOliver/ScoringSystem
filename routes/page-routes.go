package routes

import "github.com/gin-gonic/gin"

type PageController interface {
	HomePage(ctx *gin.Context)
	EventsPage(ctx *gin.Context)
	EventPage(ctx *gin.Context)
	ScoreHomePage(ctx *gin.Context)
	ScorePage(ctx *gin.Context)
}

func Pages(pages *gin.RouterGroup, pageController PageController) {
	pages.GET("/", pageController.HomePage)
	// Endpoint for All Events page.
	pages.GET("/events", pageController.EventsPage)
	events := pages.Group("/events")
	{
		// Endpoint for individual event pages.
		events.GET("/:eventId", pageController.EventPage)
	}
	// Endpoint for score page
	pages.GET("/score", pageController.ScoreHomePage)

	pages.GET("/score/:eventId", pageController.ScorePage)
}
