package countthedigit

import (
	"testing"
)

func TestNbDig(t *testing.T) {
	tables := []struct {
		x int
		y int
		z int
	}{
		{550, 5, 213},
		{5750, 0, 4700},
	}

	for _, table := range tables {
		total := NbDig(table.x, table.y)
		if total != table.z {
			t.Errorf("Testing of (%d) was incorrect, got: %d, want: %d.", table.x, total, table.z)
		}
	}
}
