package main

import (
	"strconv"
	"time"

	"github.com/Egor430-8/project/calendar"
)

func main() {
	for i := range 5 {
		calendar.AddEvents("Событие№"+strconv.Itoa(i+1), time.Now() )
	}
	calendar.ShowEvents()
}