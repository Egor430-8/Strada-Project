package events

import (
	"fmt"

	"github.com/Egor430-8/project/validation"
)

type Priority string

const (
	PriorityLow    Priority = "low"
	PriorityMedium Priority = "medium"
	PriorityHigh   Priority = "high"
	WayToValidate           = "events/priority.go/Validate: %w"
)

func (p Priority) Validate() error {
	switch p {
	case PriorityLow, PriorityMedium, PriorityHigh:
		return nil
	default:
		return fmt.Errorf(WayToValidate, validation.IncorrectPriorityError) // Done
	}
}
