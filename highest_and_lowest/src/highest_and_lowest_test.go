package highestandlowest

import (
	"testing"
)

func TestHighAndLow(t *testing.T) {
	tables := []struct {
		x string
		y string
	}{
		{"8 3 -5 42 -1 0 0 -9 4 7 4 -4", "42 -9"},
		{"8 2 3 4 -13 44 9 -3 33 9 233", "233 -13"},
		{"120 -98 39 334 382 55 8 3 884 984 646 44 -1651", "984 -1651"},
	}

	for _, table := range tables {
		total := HighAndLow(table.x)
		if total != table.y {
			t.Errorf("Testing of (%s) was incorrect, got: %s, want: %s.", table.x, total, table.y)
		}
	}
}
