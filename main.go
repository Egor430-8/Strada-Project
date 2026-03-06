package main

import (
	"github.com/Egor430-8/project/calendar"
	"github.com/Egor430-8/project/errorHandling"
	"github.com/Egor430-8/project/events"
)

func main() {
	event, err := events.NewEvent("Бег", "2026/03/13")
	errorhandling.ErrorHandling(err)
	err = calendar.AddEvents("event1", event)
	errorhandling.ErrorHandling(err)
	event, err = events.NewEvent("Теннис", "2026/05/24")
	errorhandling.ErrorHandling(err)
	err = calendar.AddEvents("event2", event)
	errorhandling.ErrorHandling(err)
	err = calendar.ShowEvents()
	errorhandling.ErrorHandling(err)
}
