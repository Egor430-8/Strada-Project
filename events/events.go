package events

import (
	"errors"
	"time"

	"github.com/Egor430-8/project/validation"
	"github.com/araddon/dateparse"
)

type Event struct {
	Title   string
	StartAt time.Time
}

func NewEvent(title string, dateStr string) (Event, error) {
	if ok := validation.IsValidTitle(title); !ok {
		return Event{}, errors.New("Заголовок введён некорректно!")
	}
	time, err := dateparse.ParseAny(dateStr)
	if err != nil {
		return Event{}, errors.New("Неверный формат даты!")
	}
	return Event{
		Title:   title,
		StartAt: time,
	}, nil
}
