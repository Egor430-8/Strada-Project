package main

import (
	"fmt"

	"github.com/Egor430-8/project/calendar"
	"github.com/Egor430-8/project/cmd"
	"github.com/Egor430-8/project/logger"
	"github.com/Egor430-8/project/storage"
)

func main() {
	logger.CreateLogger("app.log")
	defer logger.CloseFile()
	logger.System("Запуск системы")
	s := storage.NewJsonStorage("calendar.json")
	logger.System("Создано хранилище")
	c := calendar.NewCalendar(s)
	err := c.Load()
	if err != nil {
		fmt.Println(cmd.ErrorOccurred, err)
		return
	}
	logger.System("Создание и загрузка календаря")
	command := cmd.NewCmd(c)
	logger.System("Начало пользовательского ввода")
	command.Run()
}
