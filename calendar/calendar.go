package calendar

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Egor430-8/project/events"
	"github.com/Egor430-8/project/logger"
	"github.com/Egor430-8/project/storage"
	"github.com/Egor430-8/project/validation"
	"github.com/araddon/dateparse"
)

type Calendar struct {
	Events       map[string]*events.Event
	storage      storage.Store
	Notification chan string
}

const (
	WayToAddEvent            = "calendar.go/AddEvent: %w"
	WayToDeleteEvent         = "calendar.go/DeleteEvent: %w"
	WayToUpdateEvent         = "calendar.go/UpdateEvent(fullvalidation): %w"
	WayToSetEventReminder    = "calendar.go/SetEventReminder: %w"
	WayToDeleteEventReminder = "calendar.go/DeleteEventReminder: %w"
	WayToSave                = "calendar.go/Save: %w"
	WayToShowEvents          = "calendar.go/ShowEvents: %w"
)

func NewCalendar(s storage.Store) *Calendar {
	return new(Calendar{
		Events:       make(map[string]*events.Event),
		storage:      s,
		Notification: make(chan string),
	})
}

func (C *Calendar) AddEvent(title, dateStr string, priority events.Priority) (*events.Event, error) {
	logger.Info("Запуск фукции AddEvent")
	event, err := events.NewEvent(title, dateStr, priority)
	if err != nil {
		return nil, err
	}
	for _, v := range C.Events {
		if v.Title == title && v.StartAt == event.StartAt {
			return nil, fmt.Errorf(WayToAddEvent, validation.IdenticalInformationError)
		}
	}
	C.Events[event.ID] = event
	return event, nil
}

func (C Calendar) ShowEvents() error {
	logger.Info("Запуск фукции ShowEvents")
	if len(C.Events) == 0 {
		return fmt.Errorf(WayToShowEvents, validation.EmptyListError)
	}
	for _, v := range C.Events {
		fmt.Println("--------------------------------------------------------------------------------------------------")
		fmt.Println(
			v.ID,
			"-",
			v.Title,
			"-",
			v.StartAt.Format("2006-01-02 15:04"),
			"-",
			v.Priority,
		)
		if v.Reminder != nil {
			fmt.Println(
				"Есть напоминание:",
				v.Reminder.Message,
				"-",
				"Сработает:",
				v.Reminder.At.Format("2006-01-02 15:04"),
			)
		}
		fmt.Println("--------------------------------------------------------------------------------------------------")
	}
	return nil
}

func (C *Calendar) DeleteEvent(ID string) error {
	logger.Info("Запуск фукции DeleteEvent")
	err := C.isEventExist(ID)
	if err != nil {
		return fmt.Errorf(WayToDeleteEvent, validation.EventNotExistError)
	}
	delete(C.Events, ID)
	return nil
}

func (C *Calendar) UpdateEvent(ID, title, dateStr string, priority events.Priority) error {
	logger.Info("Запуск фукции UpdateEvent")
	logger.Info("Запуск фукции fullValidation")
	err, time := C.fullValidation(ID, title, dateStr, priority)
	if err != nil {
		return fmt.Errorf(WayToUpdateEvent, err)
	}
	event := C.Events[ID]
	event.Update(title, *time, priority)
	return nil
}

func (C Calendar) isEventExist(ID string) error {
	if _, ok := C.Events[ID]; !ok {
		return validation.EventNotExistError
	}
	return nil
}

func (C Calendar) fullValidation(ID, title, dateStr string, priority events.Priority) (error, *time.Time) {
	if _, ok := C.Events[ID]; !ok {
		return validation.EventNotExistError, nil
	}
	if ok := validation.IsValidTitle(title); !ok {
		return validation.IncorrectTitleError, nil
	}
	actualTime, err := dateparse.ParseAny(dateStr)
	if err != nil {
		return validation.IncorrectDateError, nil
	}
	if err = priority.Validate(); err != nil {
		return err, nil
	}
	if C.Events[ID].Title == title &&
		C.Events[ID].StartAt == actualTime &&
		C.Events[ID].Priority == priority {
		return validation.IdenticalInformationError, nil
	}
	return nil, &actualTime
}

func (C *Calendar) SetEventReminder(ID, message, dateStr string) error {
	logger.Info("Запуск фукции SetEventReminder")
	event, ok := C.Events[ID]
	if !ok {
		return fmt.Errorf(WayToSetEventReminder, validation.EventNotExistError)
	}
	if event.Reminder != nil {
		return fmt.Errorf(WayToSetEventReminder, validation.EventAlreadyHasReminderError)
	}
	if ok := validation.IsValidTitle(message); !ok {
		return fmt.Errorf(WayToSetEventReminder, validation.IncorrectTitleError)
	}
	startAt, err := dateparse.ParseAny(dateStr)
	if err != nil {
		return fmt.Errorf(WayToSetEventReminder, validation.IncorrectDateError)
	}
	err = event.AddReminder(message, startAt, C.Notify)
	if err == nil {
		return nil
	}
	return fmt.Errorf(WayToSetEventReminder, err)
}

func (C *Calendar) DeleteEventReminder(ID string) error {
	logger.Info("Запуск фукции DeleteEventReminder")
	event, ok := C.Events[ID]
	if !ok {
		return fmt.Errorf(WayToDeleteEventReminder, validation.EventNotExistError)
	}
	event.RemoveReminder()
	return nil
}

func (c *Calendar) Save() error {
	data, err := json.Marshal(c.Events)
	if err != nil {
		return fmt.Errorf(WayToSave, err)
	}
	return c.storage.Save(data)
}

func (c *Calendar) Load() error {
	data, err := c.storage.Load()
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &c.Events)
}

func (c *Calendar) Notify(msg string) {
	c.Notification <- msg
}
