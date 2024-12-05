package utils

import "fmt"

type InputError struct {
	Day int
	Msg string
}

func (e *InputError) Error() string {
	return fmt.Sprintf("Day %d: %s", e.Day, e.Msg)
}
