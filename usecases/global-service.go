package usecases

import (
	"github.com/ParadiseOliver/ScoringSystem/entity"
)

type GlobalRepository interface {
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

type globalService struct {
	repo GlobalRepository
}

func NewGlobalService(repo GlobalRepository) *globalService {
	return &globalService{
		repo: repo,
	}
}

func (service globalService) AllDisciplines() ([]entity.Discipline, error) {
	return service.repo.AllDisciplines()
}

func (service globalService) AddDiscipline(discipline *entity.Discipline) (*entity.Discipline, error) {
	return service.repo.AddDiscipline(discipline)
}

func (service globalService) DelDiscipline(id string) error {
	return service.repo.DelDiscipline(id)
}

func (service globalService) AllCategories() ([]entity.Category, error) {
	return service.repo.AllCategories()
}

func (service globalService) AddCategory(category *entity.Category) (*entity.Category, error) {
	return service.repo.AddCategory(category)
}

func (service globalService) DelCategory(id string) error {
	return service.repo.DelCategory(id)
}

func (service globalService) AllAgeGroups() ([]entity.AgeGroup, error) {
	return service.repo.AllAgeGroups()
}

func (service globalService) AddAgeGroup(ageGroup *entity.AgeGroup) (*entity.AgeGroup, error) {
	return service.repo.AddAgeGroup(ageGroup)
}

func (service globalService) DelAgeGroup(id string) error {
	return service.repo.DelAgeGroup(id)
}

func (service globalService) AllGenders() ([]entity.Gender, error) {
	return service.repo.AllGenders()
}

func (service globalService) AddGender(gender *entity.Gender) (*entity.Gender, error) {
	return service.repo.AddGender(gender)
}

func (service globalService) DelGender(id string) error {
	return service.repo.DelGender(id)
}

func (service globalService) AllCategoryGroups() ([]entity.CategoryGroup, error) {
	return service.repo.AllCategoryGroups()
}
