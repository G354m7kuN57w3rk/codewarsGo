package makeafunctionthatdoesarithmetic

import (
	"testing"
)

func TestArithmetic(t *testing.T) {
	tables := []struct {
		x int
		y int
		z string
		a int
	}{
		{8, 2, "add", 10},
		{8, 2, "subtract", 6},
		{8, 2, "multiply", 16},
		{8, 2, "divide", 4},
	}

	for _, table := range tables {
		total := Arithmetic(table.x, table.y, table.z)
		if total != table.a {
			t.Errorf("Testing of (%d %s %d) was incorrect, got: %d, want: %d.", table.x, table.z, table.y, total, table.a)
		}
	}
}
