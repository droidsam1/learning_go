package calculator

import (
	"fmt"
	"strconv"
	"strings"
)

func parse(expr string) (int, string, int, error) {
	i, op := findOperator(expr)
	if i == -1 {
		return 0, "", 0, fmt.Errorf("no operator found in %q", expr)
	}

	left, err1 := strconv.Atoi(strings.TrimSpace(expr[:i]))
	right, err2 := strconv.Atoi(strings.TrimSpace(expr[i+1:]))

	if err1 != nil {
		return 0, "", 0, fmt.Errorf("invalid left operand in %q: %w", expr, err1)
	}
	if err2 != nil {
		return 0, "", 0, fmt.Errorf("invalid right operand in %q: %w", expr, err2)
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
