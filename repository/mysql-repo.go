package repository

import (
	"log"

	"github.com/ParadiseOliver/ScoringSystem/config"
	"github.com/ParadiseOliver/ScoringSystem/entity"
)

type repo struct{}

func NewMySQLRepository() EventRepository {
	return &repo{}
}

type EventRepository interface {
	FindAll() ([]entity.Event, error)
}

func (*repo) FindAll() ([]entity.Event, error) {

	db, err := config.Connectdb()

	if err != nil {
		log.Fatalf("Failed to create a MySQL Client: %v", err)
		return nil, err
	}

	defer db.Close()

	var events []entity.Event

	res, err := db.Query("SELECT id, name FROM events")

	if err != nil {
		log.Fatal(err)
	}

	defer res.Close()
	log.Print("1")
	for res.Next() {

		var event entity.Event

		if err = res.Scan(&event.Id, &event.Name); err != nil {
			log.Fatal(err)
		}

		events = append(events, event)
	}
	log.Print("2")
	return events, nil
}
