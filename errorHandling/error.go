package errorhandling

import "fmt"

func ErrorHandling(err error) {
	if err != nil {
		fmt.Println("Произошла ошибка:", err)
		return
	}
}
