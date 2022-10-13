package main

import "math"

func isPowerOfTwo(n int) bool {
	for i := 0; i <= 31; i++ {
		if n == int(math.Pow(2, float64(i))) {
			return true
		}
	}
	return false
}
