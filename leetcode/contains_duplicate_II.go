package main

import (
	"math"
)

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if v, ok := m[nums[i]]; ok {
			if int(math.Abs(float64(i-v))) <= k {
				return true
			} else {
				m[nums[i]] = i
			}
		} else {
			m[nums[i]] = i
		}
	}

	return false
}

// func main() {
// 	nums := []int{1,2,3,1}
// 	k := 3

// 	fmt.Println(containsNearbyDuplicate(nums, k))
// }
