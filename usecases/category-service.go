package usecases

import (
	"github.com/ParadiseOliver/ScoringSystem/entity"
)

type CategoryRepository interface {
	AllDisciplines() ([]entity.Discipline, error)
	AddDiscipline(discipline *entity.Discipline) (*entity.Discipline, error)
	DelDiscipline(id string) error
	AllCategories() ([]entity.Category, error)
	AddCategory(category *entity.Category) (*entity.Category, error)
	DelCategory(id string) error
	AllAgeGroups() ([]entity.AgeGroup, error)
	AddAgeGroup(ageGroup *entity.AgeGroup) (*entity.AgeGroup, error)
	DelAgeGroup(id string) error
	AllGenders() ([]entity.Gender, error)
	AddGender(gender *entity.Gender) (*entity.Gender, error)
	DelGender(id string) error
	AllCategoryGroups() ([]entity.CategoryGroup, error)
}

type categoryService struct {
	repo CategoryRepository
}

func NewCategoryService(repo CategoryRepository) *categoryService {
	return &categoryService{
		repo: repo,
	}
}

func (service categoryService) AllDisciplines() ([]entity.Discipline, error) {
	return service.repo.AllDisciplines()
}

func (service categoryService) AddDiscipline(discipline *entity.Discipline) (*entity.Discipline, error) {
	return service.repo.AddDiscipline(discipline)
}

func (service categoryService) DelDiscipline(id string) error {
	return service.repo.DelDiscipline(id)
}

func (service categoryService) AllCategories() ([]entity.Category, error) {
	return service.repo.AllCategories()
}

func (service categoryService) AddCategory(category *entity.Category) (*entity.Category, error) {
	return service.repo.AddCategory(category)
}

func (service categoryService) DelCategory(id string) error {
	return service.repo.DelCategory(id)
}

func (service categoryService) AllAgeGroups() ([]entity.AgeGroup, error) {
	return service.repo.AllAgeGroups()
}

func (service categoryService) AddAgeGroup(ageGroup *entity.AgeGroup) (*entity.AgeGroup, error) {
	return service.repo.AddAgeGroup(ageGroup)
}

func (service categoryService) DelAgeGroup(id string) error {
	return service.repo.DelAgeGroup(id)
}

func (service categoryService) AllGenders() ([]entity.Gender, error) {
	return service.repo.AllGenders()
}

func (service categoryService) AddGender(gender *entity.Gender) (*entity.Gender, error) {
	return service.repo.AddGender(gender)
}

func (service categoryService) DelGender(id string) error {
	return service.repo.DelGender(id)
}

func (service categoryService) AllCategoryGroups() ([]entity.CategoryGroup, error) {
	return service.repo.AllCategoryGroups()
}
