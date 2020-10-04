package breakingchocolateproblem

import "testing"

func TestBreakChocolate(t *testing.T) {
	tables := []struct {
		x int
		y int
		z int
	}{
		{5, 5, 24},
		{1, 1, 0},
		{-1, 1, 0},
		{22, 5, 109},
		{44, 345, 15179},
	}

	for _, table := range tables {
		total := BreakChocolate(table.x, table.y)
		if total != table.z {
			t.Errorf("Testing of (%d, %d) was incorrect, got: %d, want: %d.", table.x, table.y, total, table.z)
		}
	}
}
