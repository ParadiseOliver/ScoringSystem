package usecases

import "github.com/ParadiseOliver/ScoringSystem/entity"

type EventService interface {
	Create(entity.Event) entity.Event
	GetAll() []entity.Event
}

type eventService struct {
	events []entity.Event
}

func New() EventService {
	return &eventService{}
}

func (service *eventService) Create(event entity.Event) entity.Event {
	service.events = append(service.events, event)
	return event
}

func (service *eventService) GetAll() []entity.Event {
	return service.events
}
