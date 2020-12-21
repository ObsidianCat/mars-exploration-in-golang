package planet

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapGenerator(t *testing.T) {
	result := MapGenerator("5 2")
	if len(result) != 3 {
		t.Errorf("Received  %d as number of columns, expected %d", len(result), 3)
	}
	if len(result[0]) != 6 {
		t.Errorf("Received  %d as number of cells in first row, expected %d", len(result[0]), 6)
	}

	t.Run("all values are zero", func(t *testing.T) {
		assert.Equal(t, 0, result[1][1], "Not initialised")
	})
}
