package twotoone

import (
	"testing"
)

func TestTwoToOne(t *testing.T) {
	tables := []struct {
		x string
		z string
		y string
	}{
		{"aretheyhere", "yestheyarehere", "aehrsty"},
		{"loopingisfunbutdangerous", "lessdangerousthancoding", "abcdefghilnoprstu"},
		{"nerve", "endings", "deginrsv"},
		{"hollow", "visions", "hilnosvw"},
		{"obacht", "gefahr", "abcefghort"},
	}

	for _, table := range tables {
		total := TwoToOne(table.x, table.z)
		if total != table.y {
			t.Errorf("Testing of (%s) and (%s) was incorrect, got: %s, want: %s.", table.x, table.z, total, table.y)
		}
	}
}
