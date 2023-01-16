package main

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/ParadiseOliver/ScoringSystem/config"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id        string `json:"id"`
	FirstName string `json:"firstname"`
	Surname   string `json:"surname"`
	Club      string `json:"club"`
	User_type string `json:"userType"`
}

type Club struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Headcoach string `json:"headcoach"`
}

type Event struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	StartDate string `json:"startdate"`
	EndDate   string `json:"enddate"`
}

type Result struct {
	Id       string `json:"id"`
	Athlete  string `json:"athlete"`
	Club     string `json:"club"`
	Agegroup string `json:"agegroup"`
	Category string `json:"category"`
	Score    string `json:"score"`
}

func init() {
	config.LoadEnvVariables()
}

func getEvents(c *gin.Context) {
	var Events []Event

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

		var event Event

		if err = res.Scan(&event.Id, &event.Name, &event.StartDate, &event.EndDate); err != nil {
			log.Fatal(err)
		}

		Events = append(Events, event)
	}

	c.IndentedJSON(http.StatusOK, Events)
}

func eventById(c *gin.Context) {
	id := c.Param("id")
	event, err := getEventById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, event)
}

func getEventById(id string) (*Event, error) {
	var event Event

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
	var newEvent Event

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

func resultById(c *gin.Context) {
	eventId := c.Param("eventId")
	eventId = "res_" + eventId
	resId := c.Param("resultId")
	result, err := getResultById(eventId, resId)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Result not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}

func getResultById(eventId string, resultId string) (*Result, error) {

	var result Result

	db, err := config.Connectdb()

	if err != nil {
		panic(err)
	}

	sql := "SELECT id, athlete_id, club_id, agegroup, category, score FROM " + eventId + " WHERE id = '" + resultId + "'"

	if err = db.QueryRow(sql).Scan(&result.Id, &result.Athlete, &result.Club, &result.Agegroup, &result.Category, &result.Score); err != nil {
		return nil, errors.New("Event not found")
	}

	return &result, nil
}

func resultsByEventId(c *gin.Context) {
	eventId := c.Param("eventId")
	eventId = "res_" + eventId
	results, err := getResultsByEventId(eventId)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Result not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, results)
}

func getResultsByEventId(eventId string) ([]Result, error) {

	var results []Result

	db, err := config.Connectdb()

	if err != nil {
		panic(err)
	}

	sql := "SELECT id, athlete_id, club_id, agegroup, category, score FROM " + eventId
	res, err := db.Query(sql)

	if err != nil {
		return nil, errors.New("Event not found")
	}

	for res.Next() {
		var result Result
		if err = res.Scan(&result.Id, &result.Athlete, &result.Club, &result.Agegroup, &result.Category, &result.Score); err != nil {
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

	r.GET("/events", getEvents)
	r.GET("/events/:id", eventById)
	r.POST("/events", createEvent)

	r.GET("/results/:eventId", resultsByEventId)
	r.GET("/results/:eventId/:resultId", resultById)

	port, err := config.MyPort()

	if err != nil {
		log.Fatal(err)
	}
	r.Run(port)
}
