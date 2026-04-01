package cmd

import (
	"fmt"
	"sync"
	"time"

	"github.com/Egor430-8/project/logger"
)

type Info struct {
	Text         string
	Err          error
	TimeCreation time.Time
}

type Logs struct {
	logList []Info
}

var mtx sync.Mutex

func NewLogs() *Logs {
	return new(Logs)
}

func (l *Logs) log(s string, err error) {
	mtx.Lock()
	l.logList = append(l.logList, Info{
		Text:         s,
		Err:          err,
		TimeCreation: time.Now(),
	})
	mtx.Unlock()
}

func (l Logs) ShowLogs() {
	logger.Info("Запуск фукции ShowLogs")
	fmt.Println("Список логов:")
	fmt.Println(Dashes)
	for _, v := range l.logList {
		switch v.Err {
		case nil:
			fmt.Println(v.Text, "-", v.TimeCreation.Format("2006-01-02 15:04"))
		default:
			fmt.Println(v.Text, "-", v.Err, "-", v.TimeCreation.Format("2006-01-02 15:04"))
		}
	}
	fmt.Println(Dashes)
}
