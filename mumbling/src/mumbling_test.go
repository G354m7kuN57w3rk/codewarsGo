package mumbling

import (
	"testing"
)

func TestAccum(t *testing.T) {
	tables := []struct {
		x string
		y string
	}{
		{"abcd", "A-Bb-Ccc-Dddd"},
		{"RqaEzty", "R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy"},
		{"cwAt", "C-Ww-Aaa-Tttt"},
	}

	for _, table := range tables {
		total := Accum(table.x)
		if total != table.y {
			t.Errorf("Testing of (%s) was incorrect, got: %s, want: %s.", table.x, total, table.y)
		}
	}
}
