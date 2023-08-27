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

	sql := fmt.Sprintf("SELECT id, event_id, athlete_id, club_id, category_id, Total FROM result WHERE event_id = '%s'", id)
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

	sql := fmt.Sprintf("SELECT id, event_id, athlete_id, club_id, category_id, Total FROM result WHERE id='%s'", id)

	if err := repo.db.QueryRow(sql).Scan(&result.ID, &result.EventID, &result.Athlete, &result.Club, &result.CategoryGroup, &result.Score); err != nil {
		return nil, errors.New("event not found")
	}

	return &result, nil
}

func (repo *mySqlResultsRepository) ResultsByAthleteId(id string) ([]entity.Result, error) {

	var results []entity.Result

	sql := fmt.Sprintf("SELECT id, event_id, athlete_id, club_id, category_id, Total FROM result WHERE athlete_id = '%s'", id)
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

func (repo *mySqlResultsRepository) UserByUserId(id string) (*entity.User, error) {

	var user entity.User

	sql := fmt.Sprintf("SELECT id, first_name, surname, club_ID, gender FROM user WHERE id='%s'", id)

	if err := repo.db.QueryRow(sql).Scan(&user.ID, &user.FirstName, &user.Surname, &user.Club, &user.Gender); err != nil {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func (repo *mySqlResultsRepository) ScoreAthlete(eventId, athleteId int, score *entity.TriScore) (*entity.Result, error) {

	sql := fmt.Sprintf("INSERT INTO result (event_id, athlete_id, club_id, category_id, routine, E1, E2, E3, E4, HD, DD, Tof, pen, total) VALUES ('%d', '%d', '1', '1', '1', '%.2f', '%.2f', '%.2f', '%.2f', '%.2f', '%.2f', '%.2f', '%.2f', '%.2f')", eventId, athleteId, score.E1, score.E2, score.E3, score.E4, score.HD, score.DD, score.Tof, score.Pen, score.Total)
	res, err := repo.db.Exec(sql)

	if err != nil {
		return nil, err
	}

	lastId, err := res.LastInsertId()

	if err != nil {
		return nil, err
	}

	var result entity.Result

	sql = fmt.Sprintf("SELECT id, event_id, athlete_id, club_id, category_id, E1 FROM result WHERE id='%d'", int(lastId))
	if err := repo.db.QueryRow(sql).Scan(&result.ID, &result.EventID, &result.Athlete, &result.Club, &result.CategoryGroup, &result.Score); err != nil {
		return nil, errors.New("result not found")
	}

	return &result, nil
}
