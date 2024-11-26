package models

import "time"

type Event struct {
	ID       int
	Name     string `binding:"required" json:"name"`
	DateTime time.Time
	UserID   int
}

var events []Event = []Event{}

func (e Event) Save() {
	events = append(events, e) // TODO: save to database
}

func GetAllEvents() []Event {
	return events
}
