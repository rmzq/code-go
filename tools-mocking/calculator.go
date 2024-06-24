package toolsmocking

import (
	"errors"
	"tools-mocking/utils"
)

type Calculator interface {
	Count(x int, y int, op rune) int
}

type calculator struct {
	math utils.Math
}

func (c calculator) Count(x int, y int, op rune) (result int, err error) {
	// var result int

	switch op {
	case '+':
		result = c.math.Add(x, y)
	case '-':

	case '*':

	case '/':

	default:
		err = errors.New("invalid operator")
	}
	return
}

func NewCalculator(math utils.Math) Calculator {
	return calculator{math: math}
}
