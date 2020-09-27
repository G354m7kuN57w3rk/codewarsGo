package removefirstandlastcharacter

import (
	"testing"
)

func TestRemoveChar(t *testing.T) {
	tables := []struct {
		x string
		y string
	}{
		{"eloquent", "loquen"},
		{"country", "ountr"},
		{"person", "erso"},
		{"place", "lac"},
		{"god", "o"},
	}

	for _, table := range tables {
		total := RemoveChar(table.x)
		if total != table.y {
			t.Errorf("Testing of (%s) was incorrect, got: %s, want: %s.", table.x, total, table.y)
		}
	}
}
