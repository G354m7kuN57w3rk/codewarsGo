package complementarydna

import (
	"testing"
)

func TestDNAStrand(t *testing.T) {
	tables := []struct {
		x string
		y string
	}{
		{"AAAA", "TTTT"},
		{"ATTGC", "TAACG"},
		{"GTAT", "CATA"},
	}

	for _, table := range tables {
		total := DNAStrand(table.x)
		if total != table.y {
			t.Errorf("Testing of (%s) was incorrect, got: %s, want: %s.", table.x, total, table.y)
		}
	}
}
