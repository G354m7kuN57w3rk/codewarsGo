package maximumlengthdifference

import (
	"testing"
)

func TestMxDifLg(t *testing.T) {
	tables := []struct {
		x []string
		y []string
		z int
	}{
		{[]string{"hoqq", "bbllkw", "oox", "ejjuyyy", "plmiis", "xxxzgpsssa", "xxwwkktt", "znnnnfqknaz", "qqquuhii", "dvvvwz"}, []string{"cccooommaaqqoxii", "gggqaffhhh", "tttoowwwmmww"}, 13},
		{[]string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"}, []string{"bbbaaayddqbbrrrv"}, 10},
		{[]string{"xoxox", "x", "xoxox", "xoxoxxx", "xxx"}, []string{}, -1},
		{[]string{}, []string{"xxxx"}, -1},
	}

	for _, table := range tables {
		total := MxDifLg(table.x, table.y)
		if total != table.z {
			t.Errorf("Testing of (%s, %s) was incorrect, got: %d, want: %d.", table.x, table.y, total, table.z)
		}
	}
}
