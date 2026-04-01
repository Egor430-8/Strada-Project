package events

import (
	"errors"
	"fmt"
	"time"

	"github.com/Egor430-8/project/logger"
	"github.com/Egor430-8/project/reminder"
	"github.com/Egor430-8/project/validation"
	"github.com/araddon/dateparse"
	"github.com/google/uuid"
)

type Event struct {
	ID       string             `json:"ID"`
	Title    string             `json:"Title"`
	StartAt  time.Time          `json:"Time"`
	Priority Priority           `json:"Priority"`
	Reminder *reminder.Reminder `json:"Reminder"`
}

const (
	WayToNewEvent    = "events/events.go/NewEvent: %w"
	WayToAddReminder = "events/events.go/AddReminder: %w"
)

func NewEvent(title, dateStr string, p Priority) (*Event, error) {
	logger.Info("Создание нового ивента через NewEvent")
	if ok := validation.IsValidTitle(title); !ok {
		return nil, fmt.Errorf(WayToNewEvent, validation.IncorrectTitleError)
	}
	time, err := dateparse.ParseAny(dateStr)
	if err != nil {
		return nil, fmt.Errorf(WayToNewEvent, validation.IncorrectDateError)
	}
	if err = p.Validate(); err != nil {
		return nil, fmt.Errorf(WayToNewEvent, err)
	}
	return new(Event{
		ID:       getNextID(),
		Title:    title,
		StartAt:  time,
		Priority: p,
		Reminder: nil,
	}), nil
}

func getNextID() string {
	return uuid.New().String()
}

func (e *Event) Update(title string, time time.Time, priority Priority) {
	logger.Info("Изменение данных ивента через Update")
	e.StartAt = time
	e.Title = title
	e.Priority = priority
}

func (e *Event) AddReminder(message string, startAt time.Time, notify func(msg string)) error {
	logger.Info("Запуск фукции AddReminder")
	e.Reminder = reminder.NewReminder(message, startAt, notify)
	err := e.Reminder.Start(notify)
	if err == nil {
		return nil
	}
	return fmt.Errorf(WayToAddReminder, err)
}

func (e *Event) RemoveReminder() error {
	logger.Info("Запуск фукции RemoveReminder")
	if e.Reminder == nil {
		return errors.New("Нельзя удалить несозданное напоминание!")
	}
	stopped := e.Reminder.Stop()
	if stopped {
		logger.Info("Таймер остановлен до срабатывания")
	} else {
		logger.Info("Таймер уже сработал или остановлен")
	}
	e.Reminder = nil
	return nil
}
