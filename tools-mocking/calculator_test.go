package toolsmocking

import (
	"testing"
	"tools-mocking/mocks"
)

func TestAdd(t *testing.T) {
	mathMock := new(mocks.Math)
	// m := new(mocks.Math)
	c := NewCalculator(mathMock)

	// c :=NewCalculator(m)

	t.Run("Success Add", func(t *testing.T) {
		// c.On("Add", 1, 2).Return(3).Once()

	})

}
