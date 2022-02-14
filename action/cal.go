package action

import "errors"

func Add(a int, b int) int {
	return a + b
}

func Sub(a int, b int) int {
	return a - b
}

func Mul(a int, b int) int {
	return a * b
}

func Div(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("math: dividend can not be zero")
	}

	return a / b, nil
}

func Sum(a []int) int {
	result := 0
	for _, num := range a {
		result += num
	}
	return result
}
