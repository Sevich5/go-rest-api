package entity

import (
	"github.com/google/uuid"
	"time"
)

type Aggregate interface {
	AddDomainEvent(event DomainEventInterface)
	GetEvents() []RecordedEvent
}

type DomainEventInterface interface{}

type RecordedEvent struct {
	EventID uuid.UUID
	Time    time.Time
	Event   DomainEventInterface
}

type AggregateRoot struct {
	events []RecordedEvent
}

func NewAggregateRoot() *AggregateRoot {
	return &AggregateRoot{events: make([]RecordedEvent, 0)}
}

func (ar *AggregateRoot) AddDomainEvent(event DomainEventInterface) {
	recorded := &RecordedEvent{
		EventID: uuid.New(),
		Time:    time.Now(),
		Event:   event,
	}
	ar.events = append(ar.events, *recorded)
}

func (ar *AggregateRoot) GetEvents() []RecordedEvent {
	return ar.events
}
