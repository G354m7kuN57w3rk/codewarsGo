package reversestrings

import (
	"testing"
)

func TestSolution(t *testing.T) {
	tables := []struct {
		x string
		y string
	}{
		{"world", "dlrow"},
		{"country", "yrtnuoc"},
		{"person", "nosrep"},
		{"place", "ecalp"},
		{"god", "dog"},
	}

	for _, table := range tables {
		total := Solution(table.x)
		if total != table.y {
			t.Errorf("Testing of (%s) was incorrect, got: %s, want: %s.", table.x, total, table.y)
		}
	}
}
