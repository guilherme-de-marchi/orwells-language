package error

import "fmt"

type Error struct {
	Name, Description string
}

func (e *Error) String() string {
	return fmt.Sprintf("%+v", *e)
}

var (
	IncompleteSyntaxError = &Error{
		"IncompleteSyntaxError",
		"Occurs when an instruction has insufficient tokens.",
	}

	InvalidSyntaxError = &Error{
		"InvalidSyntaxError",
		"Occurs when a statement does not match any syntax model.",
	}
)
