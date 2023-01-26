package main

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/ParadiseOliver/ScoringSystem/config"
	"github.com/ParadiseOliver/ScoringSystem/controllers"
	"github.com/ParadiseOliver/ScoringSystem/entity"
	"github.com/ParadiseOliver/ScoringSystem/repository"
	"github.com/ParadiseOliver/ScoringSystem/usecases"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	//gindump "github.com/tpkeeper/gin-dump"
)

var (
	eventRepo       repository.EventRepository  = repository.NewMySQLRepository()
	eventService    usecases.EventService       = usecases.New(eventRepo)
	eventController controllers.EventController = controllers.New(eventService)
)

func init() {
	config.LoadEnvVariables()
}

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func eventById(c *gin.Context) {
	id := c.Param("eventId")
	event, err := getEventById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, event)
}

func getEventById(id string) (*entity.Event, error) {
	var event entity.Event

	db, err := config.Connectdb()

	if err != nil {
		panic(err)
	}

	if err = db.QueryRow("SELECT id, name, start_date, end_date FROM events WHERE ID = ?", id).Scan(&event.Id, &event.Name, &event.StartDate, &event.EndDate); err != nil {
		return nil, errors.New("event not found")
	}

	return &event, nil
}

func allResultsByEventId(c *gin.Context) {
	eventId := c.Param("eventId")
	eventId = "results_" + eventId
	results, err := getAllResultsByEventId(eventId)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Result not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, results)
}

func getAllResultsByEventId(eventId string) ([]entity.Result, error) {

	var results []entity.Result

	db, err := config.Connectdb()

	if err != nil {
		panic(err)
	}

	sql := "SELECT id, athlete_id, club_id, agegroup, category, score FROM " + eventId
	res, err := db.Query(sql)

	if err != nil {
		return nil, errors.New("results not found")
	}

	for res.Next() {
		var result entity.Result
		if err = res.Scan(&result.Id, &result.Athlete, &result.Club, &result.Category, &result.Score); err != nil {
			panic(err)
		}

		results = append(results, result)
	}

	return results, nil
}

func resultByResultId(c *gin.Context) {
	eventId := c.Param("eventId")
	eventId = "results_" + eventId
	resId := c.Param("resultId")
	result, err := getResultByResultId(eventId, resId)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Result not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}

func getResultByResultId(eventId string, resultId string) (*entity.Result, error) {

	var result entity.Result

	db, err := config.Connectdb()

	if err != nil {
		panic(err)
	}

	sql := "SELECT id, athlete_id, club_id, agegroup, category, score FROM " + eventId + " WHERE id = '" + resultId + "'"

	if err = db.QueryRow(sql).Scan(&result.Id, &result.Athlete, &result.Club, &result.Category, &result.Score); err != nil {
		return nil, errors.New("event not found")
	}

	return &result, nil
}

func resultsByAthleteId(c *gin.Context) {
	eventId := c.Param("eventId")
	eventId = "results_" + eventId
	athleteId := c.Param("athleteId")
	log.Print(eventId, athleteId)
	results, err := getResultsByAthleteId(eventId, athleteId)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Result not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, results)
}

func getResultsByAthleteId(eventId string, athleteId string) ([]entity.Result, error) {

	var results []entity.Result

	db, err := config.Connectdb()

	if err != nil {
		panic(err)
	}

	sql := "SELECT id, athlete_id, club_id, agegroup, category, score FROM " + eventId + " WHERE athlete_id = '" + athleteId + "'"
	res, err := db.Query(sql)

	if err != nil {
		return nil, errors.New("event not found")
	}

	for res.Next() {
		var result entity.Result
		if err = res.Scan(&result.Id, &result.Athlete, &result.Club, &result.Category, &result.Score); err != nil {
			panic(err)
		}

		results = append(results, result)
	}

	return results, nil
}

func main() {

	setupLogOutput()
	//r := gin.Default()
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
			events.GET("/", func(c *gin.Context) {
				events, err := eventController.GetAll()
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				} else {
					c.JSON(http.StatusOK, events)
				}
			})
			events.POST("/", func(c *gin.Context) {
				event, err := eventController.Create(c)
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				} else {
					c.JSON(http.StatusOK, event)
				}
			})
			events.GET("/:eventId", eventById)

			events.GET("/:eventId/results", allResultsByEventId)
			events.GET("/:eventId/results/:resultId", resultByResultId)
			events.GET("/:eventId/results/athlete/:athleteId", resultsByAthleteId)
		}
	}
	r.GET("/test", func(c *gin.Context) {
		events, err := eventController.GetAll()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, events)
		}
	})
	r.POST("/test", func(c *gin.Context) {
		event, err := eventController.Create(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, event)
		}
	})

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
	r.Run(port)
}
