package isthisatriangle

import (
	"testing"
)

func TestIsTriangle(t *testing.T) {
	tables := []struct {
		x int
		y int
		z int
		a bool
	}{
		{5, 1, 2, false},
		{1, 2, 5, false},
		{2, 5, 1, false},
		{4, 2, 3, true},
		{5, 1, 5, true},
		{2, 2, 2, true},
		{-1, 2, 3, false},
	}

	for _, table := range tables {
		total := IsTriangle(table.x, table.y, table.z)
		if total != table.a {
			t.Errorf("Testing of (%d, %d, %d) was incorrect, got: %t, want: %t.", table.x, table.y, table.z, total, table.a)
		}
	}
}
