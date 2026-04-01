package calendar

import (
	"testing"

	"github.com/Egor430-8/project/events"
	"github.com/Egor430-8/project/logger"
)

func TestIsEventExist(t *testing.T) {
	ID := "12345"
	err := Calendar.isEventExist(Calendar{}, ID)
	if err == nil {
		t.Errorf("Expected an error for ID, got none")
	}
}

func TestDeleteEvent(t *testing.T) {
	logger.CreateLogger("app.log")
	defer logger.CloseFile()
	ID := "12345"
	calendar := Calendar{
		Events: make(map[string]*events.Event),
	}
	err := calendar.DeleteEvent(ID)
	if err == nil {
		t.Errorf("Expected an error for ID, got none")
	}
}

func TestAddEvent(t *testing.T) {
	logger.CreateLogger("app.log")
	defer logger.CloseFile()
	title := "lol"
	dateStr := "2026-03-30"
	priority := events.PriorityHigh
	calendar := Calendar{
		Events: make(map[string]*events.Event),
	}
	_, err := calendar.AddEvent(title, dateStr, priority)
	if err != nil {
		t.Errorf("Ошибки не ожидалось, получено %v", err)
	}

	title = "12"
	dateStr = "2026-03-30"
	priority = events.PriorityMedium
	_, err = calendar.AddEvent(title, dateStr, priority)
	if err == nil {
		t.Errorf("Ожидалась ошибка для заголовка, получено %v", err)
	}

	title = "123"
	dateStr = "2026-13-30"
	priority = events.PriorityMedium
	_, err = calendar.AddEvent(title, dateStr, priority)
	if err == nil {
		t.Errorf("Ожидалась ошибка для даты, получено %v", err)
	}
}
