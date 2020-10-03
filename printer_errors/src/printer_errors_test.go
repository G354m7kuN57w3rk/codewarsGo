package printererrors

import (
	"testing"
)

func TestPrinterError(t *testing.T) {
	tables := []struct {
		x string
		y string
	}{
		{"aaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyz", "3/56"},
		{"kkkwwwaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyz", "6/60"},
		{"kkkwwwaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyzuuuuu", "11/65"},
	}

	for _, table := range tables {
		total := PrinterError(table.x)
		if total != table.y {
			t.Errorf("Testing of (%s) was incorrect, got: %s, want: %s.", table.x, total, table.y)
		}
	}
}
