package math

import (
	"testing"
)

func TestSum(t *testing.T) {
	sum := Sum([]int{10, -2, 3})
	if sum != 11 {
		t.Errorf("FAIL: Wanted 11 but got %d", sum)
	}
}

func TestAdd(t *testing.T) {
	add := Add(2, 2)
	if add != 4 {
		t.Errorf("FAIL: Wanted 3 but got %d", add)
	}
}
