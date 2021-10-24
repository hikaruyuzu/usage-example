package calculate

import "errors"

func Add(values []int) int {
	var result int
	for _, value := range values {
		result += value
	}
	return result
}

func Calculate(val1, val2 int, operator string) (int, error) {
	var result int
	var errorResult error
	switch operator {
	case "+":
		result = val1 + val2
	case "-":
		result = val1 - val2
	case "*":
		result = val1 * val2
	case "/":
		result = val1 / val2
	default:
		errorResult = errors.New("Invalid operator")

	}
	return result, errorResult
}
