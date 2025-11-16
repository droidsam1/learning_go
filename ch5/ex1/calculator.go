package calculator

import "errors"

var (
	ErrDivisionByZero = errors.New("division by zero")
)

func add(a, b int) (int, error) {
	return a + b, nil
}

func multiply(a, b int) (int, error) {
	return a * b, nil
}

func substraction(a, b int) (int, error) {
	return a - b, nil
}

func division(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, ErrDivisionByZero
	}
	return dividend / divisor, nil
}

func perform(exp string) (int, error) {

	operations := map[string]func(int, int) (int, error){
		"-": substraction,
		"+": add,
		"*": multiply,
		"/": division,
	}

	left, op, right, err := parse(exp)
	if err != nil {
		return 0, err
	}

	return operations[string(op)](left, right)
}
