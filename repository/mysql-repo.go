package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/ParadiseOliver/ScoringSystem/entity"
)

type eventRepository struct {
	db *sql.DB
}

// TODO: Have a look at the sqlc package. You write SQL and it generates entities and queries.

func NewMySQLRepository(db *sql.DB) *eventRepository { // TODO: Struct should be called MySQLRepository?
	return &eventRepository{
		db: db,
	}
}

func (repo *eventRepository) FindAll() ([]entity.Event, error) {

	/* 	db, err := config.Connectdb() // db (connection) should be passed to NewMySQLRepository and saved on struct, then accessed here.

	   	if err != nil {
	   		return nil, err
	   	}

	   	defer func(*sql.DB) {
	   		err := db.Close()
	   		if err != nil {
	   			log.Printf("Failed to close db connection: %v", err)
	   		}
	   	}(db) */

	var events []entity.Event

	res, err := repo.db.Query("SELECT id, name FROM events")

	if err != nil {
		return nil, err
	}

	defer res.Close()

	for res.Next() {

		var event entity.Event

		if err = res.Scan(&event.ID, &event.Name); err != nil {
			return nil, err
		}

		events = append(events, event)
	}
	return events, nil
}

func (repo *eventRepository) CreateEvent(event *entity.Event) (*entity.Event, error) {
	/*
		db, err := config.Connectdb()

		if err != nil {
			log.Fatalf("Failed to create a DB Connection: %v", err)
			return nil, err
		}

		defer db.Close() */

	sql := fmt.Sprintf("INSERT INTO events (name) VALUES ('%s')", event.Name)
	res, err := repo.db.Exec(sql)

	if err != nil {
		panic(err.Error())
	}

	lastId, err := res.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}

	event.ID = strconv.Itoa(int(lastId)) // Can use RETURNING in sql with sqlc

	return event, nil
}

func (repo *eventRepository) EventById(id string) (*entity.Event, error) {
	/* 	db, err := config.Connectdb()

	   	if err != nil {
	   		log.Fatalf("Failed to create a DB Connection: %v", err)
	   		return nil, err
	   	} */

	var event entity.Event

	if err := repo.db.QueryRow("SELECT id, name, is_private FROM events WHERE ID = ?", id).Scan(&event.ID, &event.Name, &event.IsPrivate); err != nil {
		return nil, err
	}

	return &event, nil
}

func (repo *eventRepository) AllResultsByEventId(id string) ([]entity.Result, error) {
	/*
		db, err := config.Connectdb()

		if err != nil {
			log.Fatalf("Failed to create a DB Connection: %v", err)
			return nil, err
		} */

	var results []entity.Result

	sql := fmt.Sprintf("SELECT id, athlete_id, club_id, category_id, agegroup_id, score FROM results_1 WHERE event_id = '%s'", id)
	res, err := repo.db.Query(sql)

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

func (repo *eventRepository) ResultByResultId(id string) (*entity.Result, error) {
	/*
		db, err := config.Connectdb()

		if err != nil {
			log.Fatalf("Failed to create a DB Connection: %v", err)
			return nil, err
		} */

	var result entity.Result

	sql := fmt.Sprintf("SELECT id, athlete_id, club_id, category_id, agegroup_id, score FROM results_1 WHERE id='%s'", id)

	if err := repo.db.QueryRow(sql).Scan(&result.Id, &result.Athlete, &result.Club, &result.Category, &result.Agegroup, &result.Score); err != nil {
		return nil, errors.New("event not found")
	}

	return &result, nil
}

func (repo *eventRepository) ResultsByAthleteId(id string) ([]entity.Result, error) {
	/*
		db, err := config.Connectdb()

		if err != nil {
			panic(err)
		} */

	var results []entity.Result

	sql := fmt.Sprintf("SELECT id, athlete_id, club_id, category_id, agegroup_id, score FROM results_1 WHERE athlete_id = '%s'", id)
	res, err := repo.db.Query(sql)

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
