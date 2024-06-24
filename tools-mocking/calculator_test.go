package toolsmocking

import (
	"testing"
	"tools-mocking/mocks"
)

func TestAdd(t *testing.T) {

	m := new(mocks.Math)

	// c :=NewCalculator(m)

	t.Run("Success Add", func(t *testing.T) {
		m.On("Add", 1, 2).Return(3).Once()
	})

}
