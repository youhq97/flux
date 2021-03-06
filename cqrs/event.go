package cqrs

import (
	"reflect"
	"time"

	uuid "github.com/satori/go.uuid"
)

type EventMetaData struct {
	Id               string `json:"id"`
	OccuredAt        string `json:"occured_at"`
	AggregateVersion int    `json:"aggregate_version"`
	AggregateName    string `json:"aggregate_name"`
	Type             string `json:"type"`
}

//Every action on an aggregate emits an event, which is wrapped and saved
type Event struct {
	EventMetaData
	Payload interface{} `json:"payload"`
}

func (e *EventMetaData) Deserialize(data []byte) {
	deserialize(data, e)
}

func (e *Event) Deserialize(data []byte) {
	deserialize(data, e)
}

func (e *Event) Serialize() []byte {
	return serialize(e)
}

func (e *EventMetaData) Serialize() []byte {
	return serialize(e)
}

//Create new event
func NewEvent(aggregateName string, aggregateVersion int, payload interface{}) Event {
	meta := EventMetaData{
		Id:               uuid.NewV4().String(),
		AggregateVersion: aggregateVersion,
		AggregateName:    aggregateName,
		OccuredAt:        time.Now().Format(time.ANSIC),
		Type:             reflect.TypeOf(payload).String(),
	}
	return Event{meta, payload}
}

//Makes a event object from metadata and payload
func MakeEvent(meta EventMetaData, payload interface{}) Event {
	return Event{meta, payload}
}
