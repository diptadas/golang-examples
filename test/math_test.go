package test

import "testing"

func TestSum(t *testing.T) {
	ret := Sum([]int{1, 2, 3})
	if ret != 6 {
		t.Error("Expected 6, got ", ret)
	}
}