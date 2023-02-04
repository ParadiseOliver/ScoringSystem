package usecases

import (
	"github.com/ParadiseOliver/ScoringSystem/entity"
)

type EventRepository interface {
	FindAll() ([]entity.Event, error)
	CreateEvent(event *entity.Event) (*entity.Event, error)
	EventById(id string) (*entity.Event, error)
	AllResultsByEventId(id string) ([]entity.Result, error)
	ResultByResultId(id string) (*entity.Result, error)
	ResultsByAthleteId(id string) ([]entity.Result, error)
	AllDisciplines() ([]entity.Discipline, error)
	AddDiscipline(discipline *entity.Discipline) (*entity.Discipline, error)
	DelDiscipline(id string) error
	AllCategories() ([]entity.Category, error)
	AllAgeGroups() ([]entity.AgeGroup, error)
	AllGenders() ([]entity.Gender, error)
	AllCategoryGroups() ([]entity.CategoryGroup, error)
}

type eventService struct {
	repo EventRepository
}

func New(repo EventRepository) *eventService {
	return &eventService{
		repo: repo,
	}
}

func (service eventService) CreateEvent(event *entity.Event) (*entity.Event, error) {
	return service.repo.CreateEvent(event)
}

func (service eventService) GetAll() ([]entity.Event, error) {
	return service.repo.FindAll()
}

func (service eventService) GetEventById(id string) (*entity.Event, error) {
	return service.repo.EventById(id)
}

func (service eventService) AllResultsByEventId(id string) ([]entity.Result, error) {
	return service.repo.AllResultsByEventId(id)
}

func (service eventService) ResultByResultId(id string) (*entity.Result, error) {
	return service.repo.ResultByResultId(id)
}

func (service eventService) ResultsByAthleteId(id string) ([]entity.Result, error) {
	return service.repo.ResultsByAthleteId(id)
}
