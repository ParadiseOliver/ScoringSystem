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
