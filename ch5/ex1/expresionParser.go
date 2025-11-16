package calculator

import (
	"fmt"
	"strconv"
)

func parse(expr string) (int, string, int, error) {
	i, op := findOperator(expr)
	if i == -1 {
		return 0, "", 0, fmt.Errorf("no operator found in %q", expr)
	}

	left, err1 := strconv.Atoi(expr[:i])
	right, err2 := strconv.Atoi(expr[i+1:])

	if err1 != nil || err2 != nil {
		return 0, "", 0, fmt.Errorf("invalid number in %q", expr)
	}

	return left, op, right, nil
}

func findOperator(expr string) (int, string) {
	for i, r := range expr {
		// Leading + or - belongs to the first number
		if i == 0 && (r == '+' || r == '-') {
			continue
		}

		switch r {
		case '+', '-', '*', '/':
			return i, string(r)
		}
	}
	return -1, "" // no operator found
}
