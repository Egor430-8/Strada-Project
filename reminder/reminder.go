package reminder

import (
	"fmt"
	"time"

	"github.com/Egor430-8/project/logger"
	"github.com/Egor430-8/project/validation"
)

type Reminder struct {
	Message string
	At      time.Time
	Sent    bool
	Timer   *time.Timer
}

const (
	WayToStart = "reminder/reminder.go/Start: %w"
)

func NewReminder(message string, startAt time.Time, notify func(msg string)) *Reminder {
	return new(Reminder{
		Message: message,
		At:      startAt,
		Sent:    false,
		Timer:   nil,
	})
}

func (r *Reminder) Send(notify func(msg string)) {
	logger.Info("Запуск фукции Send")
	if r.Sent {
		return
	}
	logger.Info("Чтение напоминания из канала")
	notify(r.Message)
	r.Sent = true
}

func (r *Reminder) Stop() bool {
	logger.Info("Запуск фукции Stop")
	return r.Timer.Stop()
}

func (r *Reminder) Start(notify func(msg string)) error {
	logger.Info("Запуск фукции Start")
	duration := time.Until(r.At) - time.Duration(3*time.Hour)
	if duration <= 0 {
		return fmt.Errorf(WayToStart, validation.ReminderAlreadyTriggeredError)
	}
	r.Timer = time.AfterFunc(duration, func() { r.Send(notify) })
	return nil
}
