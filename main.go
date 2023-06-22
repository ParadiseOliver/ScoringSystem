package main

import (
	"database/sql"
	"io"
	"log"
	"os"

	"github.com/ParadiseOliver/ScoringSystem/config"
	"github.com/ParadiseOliver/ScoringSystem/controllers"
	"github.com/ParadiseOliver/ScoringSystem/repository"
	"github.com/ParadiseOliver/ScoringSystem/routes"
	"github.com/ParadiseOliver/ScoringSystem/usecases"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	logger "go.uber.org/zap"
	//gindump "github.com/tpkeeper/gin-dump"
)

func main() {

	f, err := os.Create("gin.log")
	if err != nil {
		log.Fatal(err)
	}
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	logger, err := logger.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()
	//sugar := logger.Sugar()

	// Load all environment variables.
	config.LoadEnvVariables()

	// Create connection to db.
	db, err := config.Connectdb()

	if err != nil {
		logger.Sugar().Fatalf("Failed to open db connection: %v", err)
	}

	// Close connection to db.
	defer func(*sql.DB) {
		err := db.Close()
		if err != nil {

			logger.Sugar().Fatalf("Failed to close db connection: %v", err)
		}
	}(db)

	// Declare repos, services and controllers.
	eventRepo := repository.NewMySQLEventRepository(db)
	categoryRepo := repository.NewMySQLCategoryRepository(db)
	resultsRepo := repository.NewMySQLResultsRepository(db)
	eventService := usecases.NewEventService(eventRepo)
	categoryService := usecases.NewCategoryService(categoryRepo)
	resultsService := usecases.NewResultsService(resultsRepo)
	eventController := controllers.NewEventController(eventService)
	categoryController := controllers.NewCategoryController(categoryService)
	resultsController := controllers.NewResultsController(resultsService)

	// Define gin server.
	r := gin.New()

	r.Static("/css", "./templates/css")
	r.LoadHTMLGlob("templates/*.html")

	//r.Use(gin.Recovery(), gin.Logger(), gindump.Dump(), middleware.BasicAuth())
	r.Use(gin.Recovery(), gin.Logger())

	// Group of all api v1 endpoints. Should it be a group for all api and a sub group for v1?
	v1 := r.Group("/api/v1")
	{
		routes.Events(v1.Group("/events"), eventController)
		routes.Categories(v1.Group("/category"), categoryController)
		routes.Results(v1.Group("/results"), resultsController)
	}

	// Routes for html pages to display.
	pages := r.Group("/pages")
	{
		// Endpoint for All Events page.
		pages.GET("/events", eventController.EventsPage)
		events := pages.Group("/events")
		{
			// Endpoint for individual event pages.
			events.GET("/:eventId", eventController.EventPage)
		}
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not found"})
	})

	port, err := config.MyPort()
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(r.Run(port))
}
