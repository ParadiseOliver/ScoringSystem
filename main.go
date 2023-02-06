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

	config.LoadEnvVariables()

	db, err := config.Connectdb()

	if err != nil {
		log.Printf("Failed to open db connection: %v", err)
	}

	defer func(*sql.DB) {
		err := db.Close()
		if err != nil {

			log.Printf("Failed to close db connection: %v", err)
		}
	}(db)

	eventRepo := repository.NewMySQLEventRepository(db)
	categoryRepo := repository.NewMySQLCategoryRepository(db)
	resultsRepo := repository.NewMySQLResultsRepository(db)
	eventService := usecases.NewEventService(eventRepo)
	categoryService := usecases.NewCategoryService(categoryRepo)
	resultsService := usecases.NewResultsService(resultsRepo)
	eventController := controllers.NewEventController(eventService)
	categoryController := controllers.NewCategoryController(categoryService)
	resultsController := controllers.NewResultsController(resultsService)

	r := gin.New()

	r.Static("/css", "./templates/css")
	r.LoadHTMLGlob("templates/*.html")

	//r.Use(gin.Recovery(), gin.Logger(), gindump.Dump(), middleware.BasicAuth())
	r.Use(gin.Recovery(), gin.Logger())

	v1 := r.Group("/api/v1")
	{
		routes.Events(v1.Group("/events"), eventController)
		routes.Categories(v1.Group("/category"), categoryController)
		routes.Results(v1.Group("/results"), resultsController)
	}

	pages := r.Group("/pages")
	{
		pages.GET("/events", eventController.EventsPage)
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
