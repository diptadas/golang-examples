package test

import "testing"

type testPair struct {
	input  []int
	output int
}

var tests = []testPair{
	{[]int{0, 1, 2}, 3},
	{[]int{1, 2, 3}, 5}, // fail
	{[]int{2, 3, 4}, 9},
	{[]int{1, 1, 1}, 2}, // fail
	{[]int{2, 2, 2}, 6},
}

func TestSum(t *testing.T) {
	for _, item := range tests {
		ret := Sum(item.input)
		if ret != item.output {
			t.Errorf("Expected %v, got %v", item.output, ret)
		}
	}
}
