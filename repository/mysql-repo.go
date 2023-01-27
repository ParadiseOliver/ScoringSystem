package repository

import (
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
