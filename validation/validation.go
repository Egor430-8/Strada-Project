package validation

import (
	"fmt"
	"regexp"
)

func IsValidTitle(title string) bool {
	pattern := "^[а-яА-ЯёЁ0-9-a-zA-Z0]+$"
	matched, err := regexp.MatchString(pattern, title)
	if err != nil {
		return false
	}
	return matched
}

func ErrorHandling(err error) {
	if err != nil {
		fmt.Println("Произошла ошибка:", err)
		return
	}
}
