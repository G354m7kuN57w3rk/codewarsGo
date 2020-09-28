package shortestword

import (
	"testing"
)

func TestFindShort(t *testing.T) {
	tables := []struct {
		x string
		y int
	}{
		{"bitcoin take over the world maybe who knows perhaps", 3},
		{"turns out random test cases are easier than writing out basic ones", 3},
		{"There's no reason no sense no meaning Behind my awkward smile", 2},
	}

	for _, table := range tables {
		total := FindShort(table.x)
		if total != table.y {
			t.Errorf("Testing of (%s) was incorrect, got: %d, want: %d.", table.x, total, table.y)
		}
	}
}
