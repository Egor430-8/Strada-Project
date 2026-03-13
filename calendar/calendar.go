package calendar

import (
	"fmt"
	"time"

	"github.com/Egor430-8/project/events"
	"github.com/Egor430-8/project/validation"
	"github.com/araddon/dateparse"
)

var AllEvents = make(map[string]events.Event)

func AddEvents(title string, dateStr string) (string, error){
	event, err := events.NewEvent(title, dateStr)
	if err != nil {
		return "", err
	}
	if _, ok := AllEvents[event.ID]; ok {
		return "", validation.TitleAlreadyExistsError
	}
	AllEvents[event.ID] = *event
	return event.ID, nil
}

func ShowEvents() error {
	if len(AllEvents) == 0 {
		return validation.EmptyListError
	}
	for _, v := range AllEvents {
		fmt.Println(v.Title, "-", v.StartAt.Format("2006-01-02 15:04"))
	}
	return nil
}

func DeleteEvent(ID string) error {
	err := isEventExist(ID)
	if err != nil {
		return err
	}
	delete(AllEvents, ID)
	return nil
}

func EditEvent(ID, title, dateStr string) error {
	err, time := fullValidation(ID, title, dateStr)
	if err != nil {
		return err
	}
	AllEvents[ID] = events.Event{
		Title:   title,	
		StartAt: time,
	}
	return nil
}

func isEventExist(ID string) error {
	if _, ok := AllEvents[ID]; !ok {
		return validation.TitleAlreadyExistsError
	}
	return nil
}

func fullValidation(ID, title, dateStr string) (error, time.Time) {
	if _, ok := AllEvents[ID]; !ok {
		return validation.TitleNotExistError, time.Now()
	}
	if ok := validation.IsValidTitle(title); !ok {
		return validation.IncorrectTitleError, time.Now()
	}
	actualTime, err := dateparse.ParseAny(dateStr)
	if err != nil {
		return validation.IncorrectDateError, time.Now()
	}
	if AllEvents[ID].Title == title && AllEvents[ID].StartAt == actualTime {
		return validation.IdenticalInformationError, time.Now()
	}
	return nil, actualTime
}
