package planet

import (
	"testing"
)

func TestMapGenerator(t *testing.T) {
	result := MapGenerator("2 3")
	if len(result) != 2 {
		t.Errorf("Received  %d as number of columns, expected %d", len(result), 2)
	}
	if len(result[0]) != 3 {
		t.Errorf("Received  %d as number of cells in first row, expected %d", len(result[0]), 3)
	}
}
