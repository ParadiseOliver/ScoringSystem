package usecases

import (
	"github.com/ParadiseOliver/ScoringSystem/entity"
)

type EventService interface { // TODO: Move me to transport layer (currently main)
	GetAll() ([]entity.Event, error)
	CreateEvent(event *entity.Event) (*entity.Event, error)
	GetEventById(id string) (*entity.Event, error)
	AllResultsByEventId(id string) ([]entity.Result, error)
	ResultByResultId(id string) (*entity.Result, error)
	ResultsByAthleteId(id string) ([]entity.Result, error)
	AllAgeGroups() ([]entity.AgeGroup, error)
}

type EventRepository interface {
	FindAll() ([]entity.Event, error)
	CreateEvent(event *entity.Event) (*entity.Event, error)
	EventById(id string) (*entity.Event, error)
	AllResultsByEventId(id string) ([]entity.Result, error)
	ResultByResultId(id string) (*entity.Result, error)
	ResultsByAthleteId(id string) ([]entity.Result, error)
	AllAgeGroups() ([]entity.AgeGroup, error)
}

type eventService struct {
	repo EventRepository
}

func New(repo EventRepository) EventService {
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

func (service eventService) AllAgeGroups() ([]entity.AgeGroup, error) {
	return service.repo.AllAgeGroups()
}
