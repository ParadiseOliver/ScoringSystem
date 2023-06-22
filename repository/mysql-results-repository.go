package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/ParadiseOliver/ScoringSystem/entity"
)

type mySqlResultsRepository struct {
	db *sql.DB
}

// TODO: Have a look at the sqlc package. You write SQL and it generates entities and queries.

func NewMySQLResultsRepository(db *sql.DB) *mySqlResultsRepository {
	return &mySqlResultsRepository{
		db: db,
	}
}

func (repo *mySqlResultsRepository) AllResultsByEventId(id string) ([]entity.Result, error) {

	var results []entity.Result

	sql := fmt.Sprintf("SELECT id, event_id, athlete_id, club_id, category_id, score FROM results_1 WHERE event_id = '%s'", id)
	res, err := repo.db.Query(sql)

	if err != nil {
		return nil, errors.New("results not found")
	}

	for res.Next() {
		var result entity.Result
		if err = res.Scan(&result.ID, &result.EventID, &result.Athlete, &result.Club, &result.CategoryGroup, &result.Score); err != nil {
			panic(err.Error())
		}

		results = append(results, result)
	}

	return results, nil
}

func (repo *mySqlResultsRepository) ResultByResultId(id string) (*entity.Result, error) {

	var result entity.Result

	sql := fmt.Sprintf("SELECT id, event_id, athlete_id, club_id, category_id, score FROM results_1 WHERE id='%s'", id)

	if err := repo.db.QueryRow(sql).Scan(&result.ID, &result.EventID, &result.Athlete, &result.Club, &result.CategoryGroup, &result.Score); err != nil {
		return nil, errors.New("event not found")
	}

	return &result, nil
}

func (repo *mySqlResultsRepository) ResultsByAthleteId(id string) ([]entity.Result, error) {

	var results []entity.Result

	sql := fmt.Sprintf("SELECT id, event_id, athlete_id, club_id, category_id, score FROM results_1 WHERE athlete_id = '%s'", id)
	res, err := repo.db.Query(sql)

	if err != nil {
		return nil, errors.New("result not found")
	}

	for res.Next() {
		var result entity.Result
		if err = res.Scan(&result.ID, &result.EventID, &result.Athlete, &result.Club, &result.CategoryGroup, &result.Score); err != nil {
			panic(err)
		}

		results = append(results, result)
	}

	return results, nil
}
