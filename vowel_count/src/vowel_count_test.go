package vowelcount

import (
	"testing"
)

func TestGetCount(t *testing.T) {
	tables := []struct {
		x string
		y int
	}{
		{"abracadabra", 5},
		{"test", 1},
		{"mechazawa", 4},
		{"god", 1},
		{"uuoiaduuoueevv", 11},
		{"xxeieeffkokoaafduiiooggh", 13},
	}

	for _, table := range tables {
		total := GetCount(table.x)
		if total != table.y {
			t.Errorf("Testing of (%s) was incorrect, got: %d, want: %d.", table.x, total, table.y)
		}
	}
}
