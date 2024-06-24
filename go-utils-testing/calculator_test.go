package goutilstesting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTambah(t *testing.T) {
	t.Run("Skenario berhasil", func(t *testing.T) {
		result := Tambah(2, 2)
		assert.Equal(t, 4, result)
	})
	t.Run("Skenario gagal", func(t *testing.T) {
		result := Kurang(6, 2)
		assert.NotEqual(t, 4, result)
	})
	result := Tambah(2, 2)
	assert.Equal(t, 4, result)
}

func TestKurang(t *testing.T) {
	result := Kurang(6, 2)
	assert.Equal(t, 4, result)
}

func TestKali(t *testing.T) {
	result := Kali(10, 2)
	assert.Equal(t, 20, result)
}

func TestBagi(t *testing.T) {
	result := Bagi(14, 2)
	assert.Equal(t, 7, result)
}
