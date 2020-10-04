package stringendswith

import (
	"testing"
)

func TestSolution(t *testing.T) {
	tables := []struct {
		x string
		y string
		z bool
	}{
		{"", "", true},
		{" ", "", true},
		{"abc", "c", true},
		{"banana", "ana", true},
		{"a", "z", false},
		{"", "t", false},
		{"$a = $b + 1", "+1", false},
		{"    ", "   ", true},
		{"1oo", "100", false},
	}

	for _, table := range tables {
		total := solution(table.x, table.y)
		if total != table.z {
			t.Errorf("Testing of (%s, %s) was incorrect, got: %t, want: %t.", table.x, table.y, total, table.z)
		}
	}
}
