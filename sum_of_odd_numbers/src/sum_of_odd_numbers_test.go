package sumofoddnumbers

import (
	"testing"
)

func TestSummation(t *testing.T) {
	tables := []struct {
		x int
		y int
	}{
		{1, 1},
		{2, 8},
		{13, 2197},
		{19, 6859},
		{41, 68921},
		{42, 74088},
		{74, 405224},
		{86, 636056},
		{93, 804357},
		{101, 1030301},
	}

	for _, table := range tables {
		total := RowSumOddNumbers(table.x)
		if total != table.y {
			t.Errorf("Testing of (%d) was incorrect, got: %d, want: %d.", table.x, total, table.y)
		}
	}
}
