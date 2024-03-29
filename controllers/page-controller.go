package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	logger "go.uber.org/zap"
)

func init() {
	logger, err := logger.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()
	//sugar := logger.Sugar()
}

type pageController struct {
}

func NewPageController() *pageController {
	//validate := validator.New()
	//validate.RegisterValidation("is-after", validators.ValidateIsAfter)
	return &pageController{}
}

func (c *pageController) HomePage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", ctx)
}

func (c *pageController) EventsPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "allEvents.html", ctx)
}

func (c *pageController) EventPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "event.html", ctx)
}

func (c *pageController) ScoreHomePage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "scoreHome.html", ctx)
}

func (c *pageController) ScorePage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "scoreEvent.html", ctx)
}
