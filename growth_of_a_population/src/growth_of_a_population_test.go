package growthofapopulation

import (
	"testing"
)

func TestNbYear(t *testing.T) {
	tables := []struct {
		x int
		y float64
		z int
		p int
		a int
	}{
		{1500, 5, 100, 5000, 15},
		{1500000, 2.5, 10000, 2000000, 10},
		{1500000, 0.25, 1000, 2000000, 94},
		{1500000, 0.25, -1000, 2000000, 151},
	}

	for _, table := range tables {
		total := NbYear(table.x, table.y, table.z, table.p)
		if total != table.a {
			t.Errorf("Testing of (%d) was incorrect, got: %d, want: %d.", table.x, total, table.a)
		}
	}
}
