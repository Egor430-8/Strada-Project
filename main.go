package main

import (
	"github.com/Egor430-8/project/calendar"
	"github.com/Egor430-8/project/events"
	"github.com/Egor430-8/project/validation"
)

func main() {
	err := calendar.ShowEvents() 
	validation.ErrorHandling(err) //Список событий пуст!

	event, err := events.NewEvent("Бег1", "2026/03/13 13:00")
	validation.ErrorHandling(err)
	err = calendar.AddEvents("event1", event)
	validation.ErrorHandling(err)

	event, err = events.NewEvent("Теннис", "2026/05/24 16:45")
	validation.ErrorHandling(err)
	err = calendar.AddEvents("event2", event)
	validation.ErrorHandling(err)

	event, err = events.NewEvent("Плавание", "2026/09/05 11:15")
	validation.ErrorHandling(err)
	err = calendar.AddEvents("event3", event)
	validation.ErrorHandling(err)

	err = calendar.EditEvent("event1", "Бег", "2026/03/13 13:30")
	validation.ErrorHandling(err)

	err = calendar.EditEvent("event1", "Бег", "2026/03/13 13:30")
	validation.ErrorHandling(err) //Были введены идентичные данные!

	err = calendar.EditEvent("event", "...", "...") 
	validation.ErrorHandling(err) //События с таким именем не существует!

	err = calendar.EditEvent("event2", "Тен.нис", "2026/05/24 16:45") 
	validation.ErrorHandling(err) //Заголовок введён некорректно!

	err = calendar.EditEvent("event2", "Теннис", "2026/13/24 16:45") 
	validation.ErrorHandling(err) //Неверный формат даты!

	err = calendar.AddEvents("event2", event)
	validation.ErrorHandling(err) //Событие с именем event2 уже существует!

	err = calendar.DeleteEvent("event")
	validation.ErrorHandling(err) //События с таким именем не существует!

	err = calendar.DeleteEvent("event3")
	validation.ErrorHandling(err)

	event, err = events.NewEvent("Т.еннис", "2026/05/24 16:45")
	validation.ErrorHandling(err) //Заголовок введён некорректно!

	event, err = events.NewEvent("Теннис", "2026/05/32  16:45")
	validation.ErrorHandling(err) //Неверный формат даты!

	err = calendar.ShowEvents()
	validation.ErrorHandling(err)
}
