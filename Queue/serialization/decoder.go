package serialization

import (
	model "ReliableQueue/model"
	"bytes"
	"encoding/gob"
)

func EventDecoder(data []byte) (model.Event, error) {
	var event model.Event
	buffer := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buffer)
	if err := dec.Decode(&event); err != nil {
		return model.Event{}, err
	}
	return event, nil
}
