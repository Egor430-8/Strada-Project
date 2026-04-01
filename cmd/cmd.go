package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/Egor430-8/project/calendar"
	"github.com/Egor430-8/project/events"
	"github.com/Egor430-8/project/logger"
	"github.com/Egor430-8/project/validation"
	"github.com/c-bata/go-prompt"
	"github.com/google/shlex"
)

const (
	AddCmd                    = "add"
	AddFormat                 = "Формат: add \"название события\" \"дата и время\" \"приоритет\""
	AddDescription            = "Добавляет событие"
	RemoveCmd                 = "remove"
	RemoveFormat              = "Формат: remove \"ID события\""
	RemoveDescription         = "Удаляет событие"
	UpdateCmd                 = "update"
	UpdateFormat              = "Формат: update \"ID события\" \"название события\" \"дата и время\" \"приоритет\""
	UpdateDescription         = "Изменяет событие"
	ListCmd                   = "list"
	ListFormat                = "Формат: list"
	ListDescription           = "Показывает все добавленные события"
	HelpCmd                   = "help"
	HelpFormat                = "Формат: help"
	HelpDescription           = "Показывает доступный список команд"
	ExitCmd                   = "exit"
	ExitFormat                = "Формат: exit"
	ExitDescription           = "Выход из программы"
	SetReminderCmd            = "setreminder"
	SetReminderFormat         = "Формат: setreminder \"ID события\" \"сообщение\" \"дата и время\""
	SetReminderDescription    = "Добавляет напоминание событию"
	DeleteReminderCmd         = "deletereminder"
	DeleteReminderFormat      = "Формат: setreminder \"ID события\""
	DeleteReminderDescription = "Удаляет напоминание у события"
	LogCmd                    = "log"
	LogFormat                 = "Формат: log"
	LogDescription            = "Список всех логов"
	ErrorOccurred             = "Произошла ошибка:"
	IncorrectCommandFormat    = "Неверный формат команды!"
	DefaultHelper             = "Введите <help> чтобы узнать доступный список команд"
	Dashes                    = "--------------------------------------------------------"
	EventAdded                = "Событие успешно добавлено!"
	EventDeleted              = "Событие успешно удалено!"
	EventUpdated              = "Событие успешно изменено!"
	ReminderAdded             = "Напоминание успешно добавлено!"
	ReminderDeleted           = "Напоминание успешно удалено!"
	LoggerErrorWriter         = "Путь + описание ошибки: %v"
)

type Cmd struct {
	calendar *calendar.Calendar
	logs     *Logs
}

func NewCmd(c *calendar.Calendar) *Cmd {
	return new(Cmd{
		calendar: c,
		logs:     new(Logs),
	})
}

func (c *Cmd) executor(input string) {
	input = strings.TrimSpace(input)
	logger.Info(fmt.Sprintf("Пользовательский ввод: %v", input))
	if input == "" {
		fmt.Println(IncorrectCommandFormat)
		fmt.Println(DefaultHelper)
		fmt.Println("")
		c.logs.log(IncorrectCommandFormat, nil)
		c.logs.log(DefaultHelper, nil)
		logger.Info(IncorrectCommandFormat)
		logger.Info(DefaultHelper)
		return
	}
	c.logs.log(input, nil)
	parts, err := shlex.Split(input)
	if err != nil {
		fmt.Println(ErrorOccurred, err)
		fmt.Println("")
		c.logs.log(ErrorOccurred, err)
		logger.Error(fmt.Sprintf(LoggerErrorWriter, err))
		return
	}
	cmd := strings.ToLower(parts[0])

	switch cmd {

	case AddCmd:
		if len(parts) != 4 {
			fmt.Println(IncorrectCommandFormat)
			fmt.Println(AddFormat)
			fmt.Println("")
			c.logs.log(IncorrectCommandFormat, nil)
			c.logs.log(AddFormat, nil)
			logger.Info(IncorrectCommandFormat)
			logger.Info(AddFormat)
			return
		}
		title := parts[1]
		date := parts[2]
		priority := events.Priority(parts[3])
		_, err := c.calendar.AddEvent(title, date, priority)
		if err != nil {
			validation.FriendlyOutput(err)
			c.logs.log(ErrorOccurred, err)
			logger.Error(fmt.Sprintf(LoggerErrorWriter, err))
			return
		}
		fmt.Println(EventAdded)
		c.logs.log(EventAdded, nil)
		logger.Info(fmt.Sprint(EventAdded, "-", "title:", title, "date:", date, "priority:", priority))
		fmt.Println("")

	case RemoveCmd:
		if len(parts) != 2 {
			fmt.Println(IncorrectCommandFormat)
			fmt.Println(RemoveFormat)
			fmt.Println("")
			c.logs.log(IncorrectCommandFormat, nil)
			c.logs.log(RemoveFormat, nil)
			logger.Info(IncorrectCommandFormat)
			logger.Info(RemoveFormat)
			return
		}
		ID := parts[1]
		err := c.calendar.DeleteEvent(ID)
		if err != nil {
			validation.FriendlyOutput(err)
			c.logs.log(ErrorOccurred, err)
			logger.Error(fmt.Sprintf(LoggerErrorWriter, err))
			fmt.Println("")
			return
		}
		fmt.Println(EventDeleted)
		c.logs.log(EventDeleted, nil)
		logger.Info(fmt.Sprint(EventDeleted, "-", "ID:", ID))
		fmt.Println("")

	case UpdateCmd:
		if len(parts) != 5 {
			fmt.Println(IncorrectCommandFormat)
			fmt.Println(UpdateFormat)
			fmt.Println("")
			c.logs.log(IncorrectCommandFormat, nil)
			c.logs.log(UpdateFormat, nil)
			logger.Info(IncorrectCommandFormat)
			logger.Info(UpdateFormat)
			return
		}
		ID := parts[1]
		Title := parts[2]
		DateStr := parts[3]
		Priority := events.Priority(parts[4])
		err = c.calendar.UpdateEvent(ID, Title, DateStr, Priority)
		if err != nil {
			validation.FriendlyOutput(err)
			c.logs.log(ErrorOccurred, err)
			logger.Error(fmt.Sprintf(LoggerErrorWriter, err))
			fmt.Println("")
			return
		}
		fmt.Println(EventUpdated)
		c.logs.log(EventUpdated, nil)
		logger.Info(fmt.Sprint(EventDeleted, "-", "ID:", ID))
		event := c.calendar.Events[ID]
		logger.Info(fmt.Sprintln("Было:", event.Title, "-", event.StartAt, "-", event.Priority))
		logger.Info(fmt.Sprintln("Стало:", Title, "-", DateStr, "-", Priority))
		_ = event
		fmt.Println("")

	case SetReminderCmd:
		if len(parts) != 4 {
			fmt.Println(IncorrectCommandFormat)
			fmt.Println(SetReminderFormat)
			fmt.Println("")
			c.logs.log(IncorrectCommandFormat, nil)
			c.logs.log(SetReminderFormat, nil)
			logger.Info(IncorrectCommandFormat)
			logger.Info(SetReminderFormat)
			return
		}
		ID := parts[1]
		Message := parts[2]
		DateStr := parts[3]
		err := c.calendar.SetEventReminder(ID, Message, DateStr)
		if err != nil {
			validation.FriendlyOutput(err)
			c.logs.log(ErrorOccurred, err)
			logger.Error(fmt.Sprintf(LoggerErrorWriter, err))
			fmt.Println("")
			return
		}
		fmt.Println(ReminderAdded)
		c.logs.log(ReminderAdded, nil)
		logger.Info(fmt.Sprintln(ReminderAdded, "-", "ID:", ID, "Message:", Message, "DateStr:", DateStr))
		fmt.Println("")

	case DeleteReminderCmd:
		if len(parts) != 2 {
			fmt.Println(IncorrectCommandFormat)
			fmt.Println(DeleteReminderFormat)
			fmt.Println("")
			c.logs.log(IncorrectCommandFormat, nil)
			c.logs.log(DeleteReminderFormat, nil)
			logger.Info(IncorrectCommandFormat)
			logger.Info(DeleteReminderFormat)
			return
		}
		ID := parts[1]
		err := c.calendar.DeleteEventReminder(ID)
		if err != nil {
			validation.FriendlyOutput(err)
			c.logs.log(ErrorOccurred, err)
			logger.Error(fmt.Sprintf(LoggerErrorWriter, err))
			fmt.Println("")
			return
		}
		fmt.Println(ReminderDeleted)
		c.logs.log(ReminderDeleted, nil)
		logger.Info(fmt.Sprint(ReminderDeleted, "-", "ID:", ID))
		fmt.Println("")

	case ListCmd:
		if len(parts) != 1 {
			fmt.Println(IncorrectCommandFormat)
			fmt.Println(ListFormat)
			fmt.Println("")
			c.logs.log(IncorrectCommandFormat, nil)
			c.logs.log(ListFormat, nil)
			logger.Info(IncorrectCommandFormat)
			logger.Info(ListFormat)
			return
		}
		err = c.calendar.ShowEvents()
		if err != nil {
			validation.FriendlyOutput(err)
			c.logs.log(ErrorOccurred, err)
			logger.Error(fmt.Sprintf(LoggerErrorWriter, err))
			fmt.Println("")
			return
		}
		fmt.Println("")

	case LogCmd:
		if len(parts) != 1 {
			fmt.Println(IncorrectCommandFormat)
			fmt.Println(LogFormat)
			fmt.Println("")
			c.logs.log(IncorrectCommandFormat, nil)
			c.logs.log(LogFormat, nil)
			logger.Info(IncorrectCommandFormat)
			logger.Info(LogFormat)
			return
		}
		c.logs.ShowLogs()
		fmt.Println("")

	case HelpCmd:
		if len(parts) != 1 {
			fmt.Println(IncorrectCommandFormat)
			fmt.Println(HelpFormat)
			fmt.Println("")
			c.logs.log(IncorrectCommandFormat, nil)
			c.logs.log(HelpFormat, nil)
			logger.Info(IncorrectCommandFormat)
			logger.Info(HelpFormat)
			return
		}
		fmt.Println("Полный список команд:")
		fmt.Println(AddCmd, "-", AddFormat, "-", AddDescription)
		fmt.Println(RemoveCmd, "-", RemoveFormat, "-", RemoveDescription)
		fmt.Println(UpdateCmd, "-", UpdateFormat, "-", UpdateDescription)
		fmt.Println(ListCmd, "-", ListFormat, "-", ListDescription)
		fmt.Println(HelpCmd, "-", HelpFormat, "-", HelpDescription)
		fmt.Println(SetReminderCmd, "-", SetReminderFormat, "-", SetReminderDescription)
		fmt.Println(DeleteReminderCmd, "-", DeleteReminderFormat, "-", DeleteReminderDescription)
		fmt.Println(LogCmd, "-", LogFormat, "-", LogDescription)
		fmt.Println(ExitCmd, "-", ExitFormat, "-", ExitDescription)
		fmt.Println("")
		logger.Info(
			fmt.Sprint(
				fmt.Sprintln("Полный список команд:"),
				fmt.Sprintln(AddCmd, "-", AddFormat, "-", AddDescription),
				fmt.Sprintln(RemoveCmd, "-", RemoveFormat, "-", RemoveDescription),
				fmt.Sprintln(UpdateCmd, "-", UpdateFormat, "-", UpdateDescription),
				fmt.Sprintln(ListCmd, "-", ListFormat, "-", ListDescription),
				fmt.Sprintln(HelpCmd, "-", HelpFormat, "-", HelpDescription),
				fmt.Sprintln(SetReminderCmd, "-", SetReminderFormat, "-", SetReminderDescription),
				fmt.Sprintln(DeleteReminderCmd, "-", DeleteReminderFormat, "-", DeleteReminderDescription),
				fmt.Sprintln(LogCmd, "-", LogFormat, "-", LogDescription),
				fmt.Sprintln(ExitCmd, "-", ExitFormat, "-", ExitDescription),
			))

	case ExitCmd:
		if len(parts) != 1 {
			fmt.Println(IncorrectCommandFormat)
			fmt.Println(ExitFormat)
			fmt.Println("")
			c.logs.log(IncorrectCommandFormat, nil)
			c.logs.log(ExitFormat, nil)
			logger.Info(IncorrectCommandFormat)
			logger.Info(ExitFormat)
			return
		}
		logger.Info("Запуск фукции Save перед завершением программы")
		err := c.calendar.Save()
		if err != nil {
			fmt.Println(ErrorOccurred, err)
			c.logs.log(ErrorOccurred, err)
			logger.Error(fmt.Sprintf(LoggerErrorWriter, err))
			return
		}
		close(c.calendar.Notification)
		fmt.Println("До скорой встречи!!!")
		c.logs.log("До скорой встречи!!!", nil)
		fmt.Println("")
		logger.System("Работа с приложением завершена")
		os.Exit(0)

	default:
		fmt.Println("Неизвестная команда:", cmd)
		fmt.Println(DefaultHelper)
		fmt.Println("")
		c.logs.log("Неизвестная команда: "+cmd, nil)
		c.logs.log(DefaultHelper, nil)
		logger.Info(fmt.Sprint("Неизвестная команда:", cmd))
		logger.Info(DefaultHelper)
	}
}

func (c *Cmd) completer(d prompt.Document) []prompt.Suggest {
	suggestions := []prompt.Suggest{
		{Text: AddCmd, Description: AddDescription},
		{Text: UpdateCmd, Description: UpdateDescription},
		{Text: ListCmd, Description: ListDescription},
		{Text: LogCmd, Description: LogDescription},
		{Text: RemoveCmd, Description: RemoveDescription},
		{Text: HelpCmd, Description: HelpDescription},
		{Text: SetReminderCmd, Description: SetReminderDescription},
		{Text: DeleteReminderCmd, Description: DeleteReminderDescription},
		{Text: ExitCmd, Description: ExitDescription},
	}
	return prompt.FilterHasPrefix(suggestions, d.GetWordBeforeCursor(), true)
}

func (c *Cmd) Run() {
	go func() {
		for msg := range c.calendar.Notification {
			fmt.Println("Напоминание:", msg)
			c.logs.log("Напоминание: "+msg, nil)
			logger.Info(fmt.Sprint("Напоминание:", msg))
		}
	}()
	p := prompt.New(
		c.executor,
		c.completer,
		prompt.OptionPrefix("> "),
	)
	p.Run()
}
