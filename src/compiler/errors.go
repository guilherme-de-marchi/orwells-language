package compiler

type Error struct {
	Name, Description string
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
