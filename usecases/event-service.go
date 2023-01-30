package usecases

import (
	"github.com/ParadiseOliver/ScoringSystem/entity"
)

type EventService interface { // TODO: Move me to transport layer (currently main)
	CreateEvent(*entity.Event) (*entity.Event, error)
	GetAll() ([]entity.Event, error)
	GetEventById(string) (*entity.Event, error)
	AllResultsByEventId(string) ([]entity.Result, error)
	ResultByResultId(string) (*entity.Result, error)
	ResultsByAthleteId(string) ([]entity.Result, error)
}

type EventRepository interface {
	FindAll() ([]entity.Event, error)
	CreateEvent(*entity.Event) (*entity.Event, error)
	EventById(string) (*entity.Event, error)
	AllResultsByEventId(string) ([]entity.Result, error)
	ResultByResultId(string) (*entity.Result, error)
	ResultsByAthleteId(string) ([]entity.Result, error)
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
