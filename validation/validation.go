package validation

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/Egor430-8/project/logger"
)

var (
	EventNotExistError            = errors.New("События с таким ID не существует!")
	EmptyListError                = errors.New("Список событий пуст!")
	IncorrectTitleError           = errors.New("Заголовок введён некорректно!")
	IncorrectDateError            = errors.New("Неверный формат даты!")
	IdenticalInformationError     = errors.New("Были введены идентичные данные!")
	DataSavingError               = errors.New("Ошибка сохранения данных!")
	DataUploadError               = errors.New("Ошибка загрузки данных!")
	IncorrectPriorityError        = errors.New("Задан неверный приоритет!")
	EventAlreadyHasReminderError  = errors.New("У данного события уже есть напоминание!")
	EmptyArchiveError             = errors.New("Архив пуст!")
	ReminderAlreadyTriggeredError = errors.New("Напоминание уже отработало или было указана неверная дата!")
	OutputForUser                 = "Вывод для пользователя:"
)

func IsValidTitle(title string) bool {
	pattern := "^[a-zA-Zа-яА-ЯёЁ0-9 ,/.]{3,200}$"
	matched, err := regexp.MatchString(pattern, title)
	if err != nil {
		return false
	}
	return matched
}

func FriendlyOutput(err error) {
	logger.Info("Запуск функции спец. вывода ошибов для пользователя")
	switch {
	case errors.Is(err, IncorrectTitleError):
		fmt.Println(
			IncorrectTitleError,
			"Заголовок должен быть длиной от 3 до 200 символов (вместе с пробелами, точками, запятыми)",
		)
		logger.Info(
			fmt.Sprintln(
				OutputForUser,
				IncorrectTitleError,
				"Заголовок должен быть длиной от 3 до 200 символов (вместе с пробелами, точками, запятыми)",
			))
		fmt.Println("")
	case errors.Is(err, IncorrectDateError):
		fmt.Println(
			IncorrectDateError,
			"Дата должна удовлетворять формату: год(полностью)-месяц-день часы:минуты",
		)
		logger.Info(
			fmt.Sprintln(
				OutputForUser,
				IncorrectDateError,
				"Дата должна удовлетворять формату: год(полностью)-месяц-день часы:минуты",
			))
		fmt.Println("")
	case errors.Is(err, EventNotExistError):
		fmt.Println(EventNotExistError)
		logger.Info(fmt.Sprintln(OutputForUser, EventNotExistError))
		fmt.Println("")
	case errors.Is(err, EmptyListError):
		fmt.Println(EmptyListError, "Сначала нужно добавить событие")
		logger.Info(fmt.Sprintln(OutputForUser, EmptyListError, "Сначала нужно добавить событие"))
		fmt.Println("")
	case errors.Is(err, IdenticalInformationError):
		fmt.Println(IdenticalInformationError)
		logger.Info(fmt.Sprintln(OutputForUser, IdenticalInformationError))
		fmt.Println("")
	case errors.Is(err, DataSavingError):
		fmt.Println(DataSavingError)
		logger.Info(fmt.Sprintln(OutputForUser, DataSavingError))
		fmt.Println("")
	case errors.Is(err, DataUploadError):
		fmt.Println(DataUploadError)
		logger.Info(fmt.Sprintln(OutputForUser, DataUploadError))
		fmt.Println("")
	case errors.Is(err, IncorrectPriorityError):
		fmt.Println(IncorrectPriorityError, "Доступны: low/medium/high")
		logger.Info(fmt.Sprintln(OutputForUser, IncorrectPriorityError, "Доступны: low/medium/high"))
		fmt.Println("")
	case errors.Is(err, EventAlreadyHasReminderError):
		fmt.Println(EventAlreadyHasReminderError)
		logger.Info(fmt.Sprintln(OutputForUser, EventAlreadyHasReminderError))
		fmt.Println("")
	case errors.Is(err, EmptyArchiveError):
		fmt.Println("Не получилось загрузить данные из zip-архива, так как", EmptyArchiveError)
		logger.Info(fmt.Sprintln(OutputForUser, "Не получилось загрузить данные из zip-архива, так как", EmptyArchiveError))
		fmt.Println("")
	case errors.Is(err, ReminderAlreadyTriggeredError):
		fmt.Println(
			ReminderAlreadyTriggeredError,
			"Необходимо проверить, чтобы дата и время напоминания были позже чем актуальная дата и время",
		)
		logger.Info(fmt.Sprintln(
			OutputForUser,
			ReminderAlreadyTriggeredError,
			"Необходимо проверить, чтобы дата и время напоминания были позже чем актуальная дата и время"))
		fmt.Println("")
	default:
		fmt.Println("Что-то пошло не так. Произошла непредсказуемая ошибка. Пожалуйста, обратитесь в поддержку приложения")
		logger.Info(fmt.Sprintln(OutputForUser, "Что-то пошло не так. Произошла непредсказуемая ошибка. Пожалуйста, обратитесь в поддержку приложения"))
		fmt.Println("")
	}
}
