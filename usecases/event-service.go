package usecases

import (
	"github.com/ParadiseOliver/ScoringSystem/entity"
	"github.com/ParadiseOliver/ScoringSystem/repository"
)

var (
	repo repository.EventRepository
)

type EventService interface {
	CreateEvent(*entity.Event) (*entity.Event, error)
	GetAll() ([]entity.Event, error)
	GetEventById(string) (*entity.Event, error)
	AllResultsByEventId(string) ([]entity.Result, error)
	ResultByResultId(string) (*entity.Result, error)
	ResultsByAthleteId(id string) ([]entity.Result, error)
}

type eventService struct{}

func New(repository repository.EventRepository) EventService {
	repo = repository
	return &eventService{}
}

func (service *eventService) CreateEvent(event *entity.Event) (*entity.Event, error) {
	return repo.CreateEvent(event)
}

func (service *eventService) GetAll() ([]entity.Event, error) {
	return repo.FindAll()
}

func (service *eventService) GetEventById(id string) (*entity.Event, error) {
	return repo.EventById(id)
}

func (service *eventService) AllResultsByEventId(id string) ([]entity.Result, error) {
	return repo.AllResultsByEventId(id)
}

func (service *eventService) ResultByResultId(id string) (*entity.Result, error) {
	return repo.ResultByResultId(id)
}

func (service *eventService) ResultsByAthleteId(id string) ([]entity.Result, error) {
	return repo.ResultsByAthleteId(id)
}
