package duplicateencoder

import (
	"testing"
)

func TestDuplicateEncode(t *testing.T) {
	tables := []struct {
		x string
		y string
	}{
		{"din", "((("},
		{"recede", "()()()"},
		{"(( @", "))(("},
		{"Success", ")())())"},
		{"Dire Dire Docks", ")))))))))))(((("},
	}

	for _, table := range tables {
		total := DuplicateEncode(table.x)
		if total != table.y {
			t.Errorf("Testing of (%s) was incorrect, got: %s, want: %s.", table.x, total, table.y)
		}
	}
}
