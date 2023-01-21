package main

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/ParadiseOliver/ScoringSystem/config"
	"github.com/ParadiseOliver/ScoringSystem/delivery"
	"github.com/ParadiseOliver/ScoringSystem/entity"
	"github.com/ParadiseOliver/ScoringSystem/usecases"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var (
	eventService    usecases.EventService    = usecases.New()
	eventController delivery.EventController = delivery.New(eventService)
)

func init() {
	config.LoadEnvVariables()
}

func getEvents(c *gin.Context) {
	var Events []entity.Event

	db, err := config.Connectdb()

	if err != nil {
		panic(err)
	}

	res, err := db.Query("SELECT * FROM events")

	if err != nil {
		log.Fatal(err)
	}

	defer res.Close()

	for res.Next() {

		var event entity.Event

		if err = res.Scan(&event.Id, &event.Name, &event.StartDate, &event.EndDate); err != nil {
			log.Fatal(err)
		}

		Events = append(Events, event)
	}

	c.IndentedJSON(http.StatusOK, Events)
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
		return nil, errors.New("Event not found")
	}

	return &event, nil
}

func createEvent(c *gin.Context) {
	var newEvent entity.Event

	db, err := config.Connectdb()

	if err != nil {
		panic(err)
	}

	if err := c.ShouldBindJSON(&newEvent); err != nil {
		log.Print(err)
		return
	}

	sql := "INSERT INTO events (name, start_date, end_date) VALUES ('" + newEvent.Name + "', '" + newEvent.StartDate + "', '" + newEvent.EndDate + "')"
	res, err := db.Exec(sql)

	if err != nil {
		panic(err.Error())
	}

	lastId, err := res.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}

	newEvent.Id = strconv.Itoa(int(lastId))

	c.IndentedJSON(http.StatusCreated, newEvent)
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

	r := gin.Default()
	//r := gin.New()
	//r.Use(gin.Recovery(), gin.Logger(), middleware.BasicAuth())

	v1 := r.Group("/api/v1")
	{
		events := v1.Group("/events")
		{
			events.GET("/", getEvents)
			events.POST("/", createEvent)
			events.GET("/:eventId", eventById)

			events.GET("/:eventId/results", allResultsByEventId)
			events.GET("/:eventId/results/:resultId", resultByResultId)
			events.GET("/:eventId/results/athlete/:athleteId", resultsByAthleteId)
		}
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
