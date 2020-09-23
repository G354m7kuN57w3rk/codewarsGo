package grasshoppersummation

import (
	"testing"
)

func TestSummation(t *testing.T) {
	tables := []struct {
		x int
		y int
	}{
		{1, 1},
		{8, 36},
		{22, 253},
		{100, 5050},
		{213, 22791},
	}

	for _, table := range tables {
		total := Summation(table.x)
		if total != table.y {
			t.Errorf("Testing of (%d) was incorrect, got: %d, want: %d.", table.x, total, table.y)
		}
	}
}
