package usecases

import (
	"github.com/ParadiseOliver/ScoringSystem/entity"
	"github.com/ParadiseOliver/ScoringSystem/repository"
)

var (
	repo repository.EventRepository
)

type EventService interface {
	Create(entity.Event) entity.Event
	GetAll() ([]entity.Event, error)
}

type eventService struct {
	events []entity.Event
}

func New(repository repository.EventRepository) EventService {
	repo = repository
	return &eventService{}
}

func (service *eventService) Create(event entity.Event) entity.Event {
	service.events = append(service.events, event)
	return event
}

func (service *eventService) GetAll() ([]entity.Event, error) {
	return repo.FindAll()
}
