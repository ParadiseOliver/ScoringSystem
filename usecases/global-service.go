package usecases

import (
	"github.com/ParadiseOliver/ScoringSystem/entity"
)

type GlobalRepository interface {
	AllDisciplines() ([]entity.Discipline, error)
	AddDiscipline(discipline *entity.Discipline) (*entity.Discipline, error)
	DelDiscipline(id string) error
	AllCategories() ([]entity.Category, error)
	AllAgeGroups() ([]entity.AgeGroup, error)
	AllGenders() ([]entity.Gender, error)
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

func (service globalService) AllAgeGroups() ([]entity.AgeGroup, error) {
	return service.repo.AllAgeGroups()
}

func (service globalService) AllGenders() ([]entity.Gender, error) {
	return service.repo.AllGenders()
}

func (service globalService) AllCategoryGroups() ([]entity.CategoryGroup, error) {
	return service.repo.AllCategoryGroups()
}
