package main

import "testing"

var nums []int = []int{3,0,1}

func TestMissingNumber(t *testing.T) {
	res := missingNumber(nums)
	if res != 2 {
		t.Errorf("expected %v, got %v", 2, res)
	}
}

func BenchmarkMissingNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		missingNumber(nums)
	}
}