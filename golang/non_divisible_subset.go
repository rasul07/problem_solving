package main

import "fmt"

func nonDivisibleSubset(k int32, s []int32) int32 {
	// Write your code here
	modulus := make(map[int32]int32)
	var result int32

	for _, v := range s {
		modulus[v%k]++
	}
	fmt.Println(modulus)
	for i, v := range modulus {

		if modulus[k-i] > v {
			if modulus[k-i] == 0 {
				result = 1
			}
			result += modulus[i-k]
		} else {
			result += v
		}
	}
	return result
}

// func main() {
// 	var (
// 		k int32 = 1
// 		s       = []int32{1, 2, 3, 4, 5}
// 	)

// 	fmt.Println(nonDivisibleSubset(k, s))
// }
