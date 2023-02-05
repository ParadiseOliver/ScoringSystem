package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/ParadiseOliver/ScoringSystem/entity"
)

type mySqlEventRepository struct {
	db *sql.DB
}

// TODO: Have a look at the sqlc package. You write SQL and it generates entities and queries.

func NewMySQLEventRepository(db *sql.DB) *mySqlEventRepository { // TODO: Struct should be called MySQLRepository?
	return &mySqlEventRepository{
		db: db,
	}
}

func (repo *mySqlEventRepository) FindAll() ([]entity.Event, error) {

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

func (repo *mySqlEventRepository) CreateEvent(event *entity.Event) (*entity.Event, error) {

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

func (repo *mySqlEventRepository) EventById(id string) (*entity.Event, error) {

	var event entity.Event

	if err := repo.db.QueryRow("SELECT id, name, is_private FROM events WHERE ID = ?", id).Scan(&event.ID, &event.Name, &event.IsPrivate); err != nil {
		return nil, err
	}

	return &event, nil
}

func (repo *mySqlEventRepository) AllResultsByEventId(id string) ([]entity.Result, error) {

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

func (repo *mySqlEventRepository) ResultByResultId(id string) (*entity.Result, error) {

	var result entity.Result

	sql := fmt.Sprintf("SELECT id, athlete_id, club_id, category_id, agegroup_id, score FROM results_1 WHERE id='%s'", id)

	if err := repo.db.QueryRow(sql).Scan(&result.Id, &result.Athlete, &result.Club, &result.Category, &result.Agegroup, &result.Score); err != nil {
		return nil, errors.New("event not found")
	}

	return &result, nil
}

func (repo *mySqlEventRepository) ResultsByAthleteId(id string) ([]entity.Result, error) {

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
