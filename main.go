package main

import (
	"database/sql"
	"io"
	"log"
	"os"

	"github.com/ParadiseOliver/ScoringSystem/config"
	"github.com/ParadiseOliver/ScoringSystem/controllers"
	"github.com/ParadiseOliver/ScoringSystem/repository"
	"github.com/ParadiseOliver/ScoringSystem/usecases"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	//gindump "github.com/tpkeeper/gin-dump"
)

func main() {

	f, err := os.Create("gin.log")
	if err != nil {
		log.Fatal(err)
	}
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	config.LoadEnvVariables()

	db, err := config.Connectdb() // db (connection) should be passed to NewMySQLRepository and saved on struct, then accessed here.

	if err != nil {
		log.Printf("Failed to open db connection: %v", err)
	}

	defer func(*sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("Failed to close db connection: %v", err)
		}
	}(db)

	eventRepo := repository.NewMySQLRepository(db)
	eventService := usecases.New(eventRepo)
	eventController := controllers.New(eventService)

	r := gin.New()
	//r.Use(gin.Recovery(), gin.Logger(), middleware.BasicAuth())
	//r.Use(gin.Recovery(), gin.Logger(), gindump.Dump())

	r.Static("/css", "./templates/css")
	r.LoadHTMLGlob("templates/*.html")

	r.Use(gin.Recovery(), gin.Logger())

	v1 := r.Group("/api/v1")
	{
		events := v1.Group("/events")
		{
			events.GET("/", eventController.GetAll)
			events.POST("/", eventController.CreateEvent)
			events.GET("/:eventId", eventController.GetEventById)

			events.GET("/result/:resultId", eventController.ResultByResultId)
			events.GET("/results/:eventId", eventController.AllResultsByEventId)
			events.GET("/athlete/:athleteId/results", eventController.ResultsByAthleteId)
		}
	}

	pages := r.Group("/pages")
	{
		pages.GET("/events", eventController.AllEvents)
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
