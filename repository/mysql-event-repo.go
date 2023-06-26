package repository

import (
	"database/sql"
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

	res, err := repo.db.Query("SELECT id, name, start_date, end_date, is_private FROM events")

	if err != nil {
		return nil, err
	}

	defer res.Close()

	for res.Next() {

		var event entity.Event

		if err = res.Scan(&event.ID, &event.Name, &event.StartDate, &event.EndDate, &event.IsPrivate); err != nil {
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

	if err := repo.db.QueryRow("SELECT id, name, start_date, end_date, is_private FROM events WHERE ID = ?", id).Scan(&event.ID, &event.Name, &event.StartDate, &event.EndDate, &event.IsPrivate); err != nil {
		return nil, err
	}

	return &event, nil
}
