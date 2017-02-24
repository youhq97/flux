package cqrs

import (
	"bytes"

	"encoding/gob"

	"time"

	"github.com/golang/glog"
	uuid "github.com/satori/go.uuid"
)

//Every action on an aggregate emits an event, which is wrapped and saved
type Event struct {
	Id               string      // Event Id
	Payload          interface{} // Event payload
	OccuredAt        string      // The time the event occured
	AggregateVersion int         // The aggregate version on which this event occurred
	AggregateName    string      // Event aggregate name
}

//Deserialize the event
func (e *Event) Deserialize(data []byte) {
	b := bytes.Buffer{}
	b.Write(data)
	d := gob.NewDecoder(&b)
	err := d.Decode(e)
	if err != nil {
		glog.Errorf("could not deserialize event %v", e)
	}
}

//Serialize the event
func (e *Event) Serialize() []byte {
	buffer := bytes.Buffer{}
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(e)
	if err != nil {
		glog.Errorf("could not serialize event %v", e)
	}
	return buffer.Bytes()
}

//Create new event
func NewEvent(aggregateName string, aggregateVersion int, payload interface{}) Event {
	gob.Register(payload)
	return Event{
		Id:               uuid.NewV4().String(),
		Payload:          payload,
		AggregateVersion: aggregateVersion,
		AggregateName:    aggregateName,
		OccuredAt:        time.Now().Format(time.ANSIC),
	}
}