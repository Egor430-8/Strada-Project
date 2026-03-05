package calendar

import (
	"fmt"
	"time"

	"github.com/Egor430-8/project/events"
	"github.com/k0kubun/pp"
)

var AllEvents = make(map[string]events.Event)

func AddEvents(title string, time time.Time) {
	AllEvents[title] = events.Event{
		Title: title,
		StartAt: time,
	}
}

func ShowEvents() error {
	if len(AllEvents) == 0 {
		return fmt.Errorf("Список событий пуст!")
	}
	for _, v := range AllEvents {
		pp.Println(v)
	}
	return nil
}