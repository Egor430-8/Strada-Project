package validation

import (
	"errors"
	"regexp"
)

var (
	TitleNotExistError        = errors.New("Событие с таким именем не существует!")
	EmptyListError            = errors.New("Список событий пуст!")
	IncorrectTitleError       = errors.New("Заголовок введён некорректно!")
	IncorrectDateError        = errors.New("Неверный формат даты!")
	IdenticalInformationError = errors.New("Были введены идентичные данные!")
	TitleAlreadyExistsError   = errors.New("Событие с таким именем уже существует!")
)

func IsValidTitle(title string) bool {
	pattern := "^[а-яА-ЯёЁ0-9-a-zA-Z0]+$"
	matched, err := regexp.MatchString(pattern, title)
	if err != nil {
		return false
	}
	return matched
}
