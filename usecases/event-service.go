package usecases

import (
	"github.com/ParadiseOliver/ScoringSystem/entity"
)

type EventRepository interface {
	FindAll() ([]entity.Event, error)
	CreateEvent(event *entity.Event) (*entity.Event, error)
	EventById(id string) (*entity.Event, error)
}

type eventService struct {
	repo EventRepository
}

func NewEventService(repo EventRepository) *eventService {
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
