package usecases

import (
	"github.com/ParadiseOliver/ScoringSystem/entity"
)

type CategoryRepository interface {
	AllDisciplines() ([]entity.Discipline, error)
	Discipline(id string) (*entity.Discipline, error)
	AddDiscipline(discipline *entity.Discipline) (*entity.Discipline, error)
	UpdateDiscipline(discipline *entity.Discipline) (*entity.Discipline, error)
	DelDiscipline(id string) error
	AllCategories() ([]entity.Category, error)
	Category(id string) (*entity.Category, error)
	AddCategory(category *entity.Category) (*entity.Category, error)
	DelCategory(id string) error
	AllAgeGroups() ([]entity.AgeGroup, error)
	AgeGroup(id string) (*entity.AgeGroup, error)
	AddAgeGroup(ageGroup *entity.AgeGroup) (*entity.AgeGroup, error)
	DelAgeGroup(id string) error
	AllGenders() ([]entity.Gender, error)
	Gender(id string) (*entity.Gender, error)
	AddGender(gender *entity.Gender) (*entity.Gender, error)
	DelGender(id string) error
	AllCategoryGroups() ([]entity.CategoryGroup, error)
	AddCategoryGroup(catGroup *entity.CategoryGroup) (*entity.CategoryGroup, error)
	CategoryGroup(id string) (*entity.CategoryGroup, error)
	DelCategoryGroup(id string) error
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

func (service categoryService) Discipline(id string) (*entity.Discipline, error) {
	return service.repo.Discipline(id)
}

func (service categoryService) AddDiscipline(discipline *entity.Discipline) (*entity.Discipline, error) {
	return service.repo.AddDiscipline(discipline)
}

func (service categoryService) UpdateDiscipline(discipline *entity.Discipline) (*entity.Discipline, error) {
	return service.repo.UpdateDiscipline(discipline)
}

func (service categoryService) DelDiscipline(id string) error {
	return service.repo.DelDiscipline(id)
}

func (service categoryService) AllCategories() ([]entity.Category, error) {
	return service.repo.AllCategories()
}

func (service categoryService) Category(id string) (*entity.Category, error) {
	return service.repo.Category(id)
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

func (service categoryService) AgeGroup(id string) (*entity.AgeGroup, error) {
	return service.repo.AgeGroup(id)
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

func (service categoryService) Gender(id string) (*entity.Gender, error) {
	return service.repo.Gender(id)
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

func (service categoryService) CategoryGroup(id string) (*entity.CategoryGroup, error) {
	return service.repo.CategoryGroup(id)
}

func (service categoryService) DelCategoryGroup(id string) error {
	return service.repo.DelCategoryGroup(id)
}

func (service categoryService) AddCategoryGroup(catGroup *entity.CategoryGroup) (*entity.CategoryGroup, error) {
	return service.repo.AddCategoryGroup(catGroup)
}
