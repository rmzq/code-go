package utils

type Math interface {
	Add(a, b int) int
}

type math struct{}

func NewMath() Math {
	return math{}
}

func (m math) Add(a, b int) int {
	return a + b
}
