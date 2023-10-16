package Output

import "ReliableQueue/model"

type EventStatus struct {
	Base
	Event model.Event
}

func (cmd *EventStatus) SetErr(e error) {
	cmd.Error = e
}

func (cmd *EventStatus) Err() error {
	return cmd.Error
}

func (cmd *EventStatus) SetVal(event model.Event) {
	cmd.Event = event
}

func (cmd *EventStatus) GetEvent() model.Event {
	return cmd.Event
}

func (cmd *EventStatus) Result() (model.Event, error) {
	return cmd.Event, cmd.Error
}
