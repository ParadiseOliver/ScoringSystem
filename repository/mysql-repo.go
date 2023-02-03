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

func NewMySQLRepository(db *sql.DB) *eventRepository { // TODO: Struct should be called MySQLRepository? Done
	return &eventRepository{
		db: db,
	}
}

func (repo *eventRepository) FindAll() ([]entity.Event, error) {

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

	var event entity.Event

	if err := repo.db.QueryRow("SELECT id, name, is_private FROM events WHERE ID = ?", id).Scan(&event.ID, &event.Name, &event.IsPrivate); err != nil {
		return nil, err
	}

	return &event, nil
}

func (repo *eventRepository) AllResultsByEventId(id string) ([]entity.Result, error) {

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

	var result entity.Result

	sql := fmt.Sprintf("SELECT id, athlete_id, club_id, category_id, agegroup_id, score FROM results_1 WHERE id='%s'", id)

	if err := repo.db.QueryRow(sql).Scan(&result.Id, &result.Athlete, &result.Club, &result.Category, &result.Agegroup, &result.Score); err != nil {
		return nil, errors.New("event not found")
	}

	return &result, nil
}

func (repo *eventRepository) ResultsByAthleteId(id string) ([]entity.Result, error) {

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

func (repo *eventRepository) AllDisciplines() ([]entity.Discipline, error) {

	res, err := repo.db.Query("SELECT disciplines_id, discipline FROM disciplines")

	if err != nil {
		return nil, err
	}

	var disciplines []entity.Discipline

	for res.Next() {
		var discipline entity.Discipline
		if err = res.Scan(&discipline.ID, &discipline.Discipline); err != nil {
			return nil, err
		}

		disciplines = append(disciplines, discipline)
	}

	return disciplines, nil
}

func (repo *eventRepository) AddDiscipline(discipline *entity.Discipline) (*entity.Discipline, error) {

	sql := fmt.Sprintf("INSERT INTO disciplines (discipline) VALUES ('%s')", discipline.Discipline)
	res, err := repo.db.Exec(sql)

	if err != nil {
		return nil, err
	}

	lastId, err := res.LastInsertId()

	if err != nil {
		return nil, err
	}

	discipline.ID = strconv.Itoa(int(lastId)) // Can use RETURNING in sql with sqlc

	return discipline, nil
}

func (repo *eventRepository) DelDiscipline(id string) error {

	sql := fmt.Sprintf("DELETE FROM disciplines WHERE disciplines_id = '%s'", id)
	_, err := repo.db.Exec(sql)

	if err != nil {
		return err
	}

	return nil
}

func (repo *eventRepository) AllCategories() ([]entity.Category, error) {

	res, err := repo.db.Query("SELECT categories_id, category FROM categories")

	if err != nil {
		return nil, err
	}

	var categories []entity.Category

	for res.Next() {
		var category entity.Category
		if err = res.Scan(&category.ID, &category.Category); err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func (repo *eventRepository) AllAgeGroups() ([]entity.AgeGroup, error) {

	res, err := repo.db.Query("SELECT agegroup_id, min_age, max_age, group_name FROM agegroups")

	if err != nil {
		return nil, err
	}

	var ageGroups []entity.AgeGroup

	for res.Next() {
		var ageGroup entity.AgeGroup
		if err = res.Scan(&ageGroup.ID, &ageGroup.MinAge, &ageGroup.MaxAge, &ageGroup.CategoryName); err != nil {
			return nil, err
		}

		ageGroups = append(ageGroups, ageGroup)
	}

	return ageGroups, nil
}

func (repo *eventRepository) AllGenders() ([]entity.Gender, error) {

	res, err := repo.db.Query("SELECT genders_id, gender FROM genders")

	if err != nil {
		return nil, err
	}

	var genders []entity.Gender

	for res.Next() {
		var gender entity.Gender
		if err = res.Scan(&gender.ID, &gender.Gender); err != nil {
			return nil, err
		}

		genders = append(genders, gender)
	}

	return genders, nil
}

func (repo *eventRepository) AllCategoryGroups() ([]entity.CategoryGroup, error) {

	res, err := repo.db.Query("SELECT cat_id, discipline_id, category_id, agegroup_id, gender_id FROM category_groups")

	if err != nil {
		return nil, err
	}

	var categoryGroups []entity.CategoryGroup

	for res.Next() {
		var categoryGroup entity.CategoryGroup
		if err = res.Scan(&categoryGroup.ID, &categoryGroup.DisciplineId, &categoryGroup.CategoryId, &categoryGroup.AgegroupId, &categoryGroup.GenderId); err != nil {
			return nil, err
		}

		categoryGroups = append(categoryGroups, categoryGroup)
	}

	return categoryGroups, nil
}
