package main

import (
	"testing"
)

func TestInverseCaptchaPart1(t *testing.T) {
	tables := []struct {
		x string
		y int
		n int
	}{
		{"1122", 1, 3},
		{"1111", 1, 4},
		{"1234", 1, 0},
		{"91212129", 1, 9},
	}

	for _, table := range tables {
		result := inverseCaptcha(table.x, table.y)
		if result != table.n {
			t.Errorf("InverseCaptcha(%s, %d) was incorrect, got: %d, want: %d.", table.x, table.y, result, table.n)
		}
	}
}
