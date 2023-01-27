package repository

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/ParadiseOliver/ScoringSystem/config"
	"github.com/ParadiseOliver/ScoringSystem/entity"
)

type EventRepository interface {
	FindAll() ([]entity.Event, error)
	CreateEvent(*entity.Event) (*entity.Event, error)
	EventById(string) (*entity.Event, error)
	AllResultsByEventId(string) ([]entity.Result, error)
	ResultByResultId(string) (*entity.Result, error)
	ResultsByAthleteId(string) ([]entity.Result, error)
}

type repo struct{}

func NewMySQLRepository() EventRepository {
	return &repo{}
}

func (*repo) FindAll() ([]entity.Event, error) {

	db, err := config.Connectdb()

	if err != nil {
		log.Fatalf("Failed to create a DB Connection: %v", err)
		return nil, err
	}

	defer db.Close()

	var events []entity.Event

	res, err := db.Query("SELECT id, name FROM events")

	if err != nil {
		log.Fatal(err)
	}

	defer res.Close()

	for res.Next() {

		var event entity.Event

		if err = res.Scan(&event.Id, &event.Name); err != nil {
			log.Fatal(err)
		}

		events = append(events, event)
	}
	return events, nil
}

func (*repo) CreateEvent(event *entity.Event) (*entity.Event, error) {

	db, err := config.Connectdb()

	if err != nil {
		log.Fatalf("Failed to create a DB Connection: %v", err)
		return nil, err
	}

	defer db.Close()

	sql := fmt.Sprintf("INSERT INTO events (name) VALUES ('%s')", event.Name)
	res, err := db.Exec(sql)

	if err != nil {
		panic(err.Error())
	}

	lastId, err := res.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}

	event.Id = strconv.Itoa(int(lastId))

	return event, nil
}

func (*repo) EventById(id string) (*entity.Event, error) {
	db, err := config.Connectdb()

	if err != nil {
		log.Fatalf("Failed to create a DB Connection: %v", err)
		return nil, err
	}

	var event entity.Event

	if err = db.QueryRow("SELECT id, name, is_private FROM events WHERE ID = ?", id).Scan(&event.Id, &event.Name, &event.IsPrivate); err != nil {
		return nil, err
	}

	return &event, nil
}

func (*repo) AllResultsByEventId(id string) ([]entity.Result, error) {

	db, err := config.Connectdb()

	if err != nil {
		log.Fatalf("Failed to create a DB Connection: %v", err)
		return nil, err
	}

	var results []entity.Result

	sql := fmt.Sprintf("SELECT id, athlete_id, club_id, category_id, agegroup_id, score FROM results_1 WHERE event_id = '%s'", id)
	res, err := db.Query(sql)

	if err != nil {
		return nil, errors.New("results not found")
	}

	for res.Next() {
		var result entity.Result
		if err = res.Scan(&result.Id, &result.Athlete, &result.Club, &result.Category, &result.Agegroup, &result.Score); err != nil {
			panic(err.Error())
		}

		results = append(results, result)
	}

	return results, nil
}

func (*repo) ResultByResultId(id string) (*entity.Result, error) {

	db, err := config.Connectdb()

	if err != nil {
		log.Fatalf("Failed to create a DB Connection: %v", err)
		return nil, err
	}

	var result entity.Result

	sql := fmt.Sprintf("SELECT id, athlete_id, club_id, category_id, agegroup_id, score FROM results_1 WHERE id='%s'", id)

	if err = db.QueryRow(sql).Scan(&result.Id, &result.Athlete, &result.Club, &result.Category, &result.Agegroup, &result.Score); err != nil {
		return nil, errors.New("event not found")
	}

	return &result, nil
}

func (*repo) ResultsByAthleteId(id string) ([]entity.Result, error) {

	db, err := config.Connectdb()

	if err != nil {
		panic(err)
	}

	var results []entity.Result

	sql := fmt.Sprintf("SELECT id, athlete_id, club_id, category_id, agegroup_id, score FROM results_1 WHERE athlete_id = '%s'", id)
	res, err := db.Query(sql)

	if err != nil {
		return nil, errors.New("event not found")
	}

	for res.Next() {
		var result entity.Result
		if err = res.Scan(&result.Id, &result.Athlete, &result.Club, &result.Category, &result.Agegroup, &result.Score); err != nil {
			panic(err)
		}

		results = append(results, result)
	}

	return results, nil
}
