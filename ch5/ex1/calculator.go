package calculator

import (
	"errors"
	"fmt"
)

var (
	ErrDivisionByZero      = errors.New("division by zero")
	ErrUnsupportedOperator = errors.New("not valid operator")
)

type operation func(int, int) (int, error)

func add(a, b int) (int, error) {
	return a + b, nil
}

func multiply(a, b int) (int, error) {
	return a * b, nil
}

func subtraction(a, b int) (int, error) {
	return a - b, nil
}

func division(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, ErrDivisionByZero
	}
	return dividend / divisor, nil
}

func Perform(exp string) (int, error) {

	operations := map[string]operation{
		"-": subtraction,
		"+": add,
		"*": multiply,
		"/": division,
	}

	left, op, right, err := parse(exp)
	if err != nil {
		return 0, err
	}

	fn, ok := operations[op]

	if !ok {
		return 0, fmt.Errorf("%w: %q", ErrUnsupportedOperator, op)
	}

	return fn(left, right)
}
