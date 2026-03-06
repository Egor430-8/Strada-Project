package calendar

import (
	"errors"
	"fmt"

	"github.com/Egor430-8/project/events"
)

var AllEvents = make(map[string]events.Event)

func AddEvents(title string, event events.Event) error {
	if _, ok := AllEvents[title]; ok {
		return errors.New("Событие с именем " + title + " уже существует!")
	}
	AllEvents[title] = event
	return nil
}

func ShowEvents() error {
	if len(AllEvents) == 0 {
		return errors.New("Список событий пуст!")
	}
	for _, v := range AllEvents {
		fmt.Println(v.Title, "-", v.StartAt.Format("2006-01-02"))
	}
	return nil
}
