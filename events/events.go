package events

import (
	"time"

	"github.com/Egor430-8/project/validation"
	"github.com/araddon/dateparse"
	"github.com/google/uuid"
)

type Event struct {
	ID      string
	Title   string
	StartAt time.Time
}

func NewEvent(title string, dateStr string) (*Event, error) {
	if ok := validation.IsValidTitle(title); !ok {
		return nil, validation.IncorrectTitleError
	}
	time, err := dateparse.ParseAny(dateStr)
	if err != nil {
		return nil, validation.IncorrectDateError
	}
	return &Event{
		ID:      getNextID(),
		Title:   title,
		StartAt: time,
	}, nil
}

func getNextID() string {
	return uuid.New().String()
}
