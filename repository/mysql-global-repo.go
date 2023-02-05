package repository

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/ParadiseOliver/ScoringSystem/entity"
)

type mySqlGlobalRepository struct {
	db *sql.DB
}

// TODO: Have a look at the sqlc package. You write SQL and it generates entities and queries.

func NewMySQLGlobalRepository(db *sql.DB) *mySqlGlobalRepository {
	return &mySqlGlobalRepository{
		db: db,
	}
}

func (repo *mySqlGlobalRepository) AllDisciplines() ([]entity.Discipline, error) {

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

func (repo *mySqlGlobalRepository) AddDiscipline(discipline *entity.Discipline) (*entity.Discipline, error) {

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

func (repo *mySqlGlobalRepository) DelDiscipline(id string) error {

	sql := fmt.Sprintf("DELETE FROM disciplines WHERE disciplines_id = '%s'", id)
	_, err := repo.db.Exec(sql)

	if err != nil {
		return err
	}

	return nil
}

func (repo *mySqlGlobalRepository) AllCategories() ([]entity.Category, error) {

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

func (repo *mySqlGlobalRepository) AddCategory(category *entity.Category) (*entity.Category, error) {

	sql := fmt.Sprintf("INSERT INTO categories (category) VALUES ('%s')", category.Category)
	res, err := repo.db.Exec(sql)

	if err != nil {
		return nil, err
	}

	lastId, err := res.LastInsertId()

	if err != nil {
		return nil, err
	}

	category.ID = strconv.Itoa(int(lastId)) // Can use RETURNING in sql with sqlc

	return category, nil
}

func (repo *mySqlGlobalRepository) DelCategory(id string) error {

	sql := fmt.Sprintf("DELETE FROM categories WHERE categories_id = '%s'", id)
	_, err := repo.db.Exec(sql)

	if err != nil {
		return err
	}

	return nil
}

func (repo *mySqlGlobalRepository) AllAgeGroups() ([]entity.AgeGroup, error) {

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

func (repo *mySqlGlobalRepository) AddAgeGroup(ageGroup *entity.AgeGroup) (*entity.AgeGroup, error) {

	sql := fmt.Sprintf("INSERT INTO agegroups (min_age, max_age, group_name) VALUES ('%d', '%d', '%s')", ageGroup.MinAge, ageGroup.MaxAge, ageGroup.CategoryName)
	res, err := repo.db.Exec(sql)

	if err != nil {
		return nil, err
	}

	lastId, err := res.LastInsertId()

	if err != nil {
		return nil, err
	}

	ageGroup.ID = strconv.Itoa(int(lastId)) // Can use RETURNING in sql with sqlc

	return ageGroup, nil
}

func (repo *mySqlGlobalRepository) DelAgeGroup(id string) error {

	sql := fmt.Sprintf("DELETE FROM agegroups WHERE agegroup_id = '%s'", id)
	_, err := repo.db.Exec(sql)

	if err != nil {
		return err
	}

	return nil
}

func (repo *mySqlGlobalRepository) AllGenders() ([]entity.Gender, error) {

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

func (repo *mySqlGlobalRepository) AddGender(gender *entity.Gender) (*entity.Gender, error) {

	sql := fmt.Sprintf("INSERT INTO genders (gender) VALUES ('%s')", gender.Gender)
	res, err := repo.db.Exec(sql)

	if err != nil {
		return nil, err
	}

	lastId, err := res.LastInsertId()

	if err != nil {
		return nil, err
	}

	gender.ID = strconv.Itoa(int(lastId)) // Can use RETURNING in sql with sqlc

	return gender, nil
}

func (repo *mySqlGlobalRepository) DelGender(id string) error {

	sql := fmt.Sprintf("DELETE FROM genders WHERE genders_id = '%s'", id)
	_, err := repo.db.Exec(sql)

	if err != nil {
		return err
	}

	return nil
}

func (repo *mySqlGlobalRepository) AllCategoryGroups() ([]entity.CategoryGroup, error) {

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
