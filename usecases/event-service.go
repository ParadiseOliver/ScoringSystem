package usecases

import (
	"github.com/ParadiseOliver/ScoringSystem/entity"
	"github.com/ParadiseOliver/ScoringSystem/repository"
)

var (
	repo repository.EventRepository
)

type EventService interface {
	Create(*entity.Event) (*entity.Event, error)
	GetAll() ([]entity.Event, error)
}

type eventService struct{}

func New(repository repository.EventRepository) EventService {
	repo = repository
	return &eventService{}
}

func (service *eventService) Create(event *entity.Event) (*entity.Event, error) {
	return repo.Save(event)
}

func (service *eventService) GetAll() ([]entity.Event, error) {
	return repo.FindAll()
}
