package utils

import "errors"

type Math interface {
	Add(a, b int) int
	Sub(a, b int) (int, error)
}

type math struct{}

func NewMath() Math {
	return math{}
}

func (m math) Add(a, b int) int {
	return a + b
}

func (m math) Sub(a, b int) (result int, err error) {
	if a < b {
		err = errors.New("a is less than b")
	}

	result = a - b

	return
}
