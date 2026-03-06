package calendar

import (
	"errors"
	"fmt"
	"time"

	"github.com/Egor430-8/project/events"
	"github.com/Egor430-8/project/validation"
	"github.com/araddon/dateparse"
)

var AllEvents = make(map[string]events.Event)
var TitleError = errors.New("События с таким именем не существует!")

func AddEvents(name string, event events.Event) error {
	if _, ok := AllEvents[name]; ok {
		return errors.New("Событие с именем " + name + " уже существует!")
	}
	AllEvents[name] = event
	return nil
}

func ShowEvents() error {
	if len(AllEvents) == 0 {
		return errors.New("Список событий пуст!")
	}
	for _, v := range AllEvents {
		fmt.Println(v.Title, "-", v.StartAt.Format("2006-01-02 15:04"))
	}
	return nil
}

func DeleteEvent(name string) error {
	err := isEventExist(name)
	if err != nil {
		return err
	}
	delete(AllEvents, name)
	return nil
}

func EditEvent(name, title, dateStr string) error {
	err, time := fullValidation(name, title, dateStr)
	if err != nil {
		return err
	}
	AllEvents[name] = events.Event{
		Title:   title,
		StartAt: time,
	}
	return nil
}

func isEventExist(name string) error {
	if _, ok := AllEvents[name]; !ok {
		return TitleError
	}
	return nil
}

func fullValidation(name, title, dateStr string) (error, time.Time) {
	if _, ok := AllEvents[name]; !ok {
		return TitleError, time.Now()
	}
	if ok := validation.IsValidTitle(title); !ok {
		return errors.New("Заголовок введён некорректно!"), time.Now()
	}
	actualTime, err := dateparse.ParseAny(dateStr)
	if err != nil {
		return errors.New("Неверный формат даты!"), time.Now()
	}
	if AllEvents[name].Title == title && AllEvents[name].StartAt == actualTime {
		return errors.New("Были введены идентичные данные!"), time.Now()
	}
	return nil, actualTime
}
