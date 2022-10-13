package main

import (
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := 0; i < n; i++ {
		nums1[m+i] = nums2[i]
	}
	sort.Ints(nums1)
}

// func main() {
// 	nums1 := []int{1, 2, 3, 0, 0, 0}
// 	m := 3
// 	nums2 := []int{2, 5, 6}
// 	n := 3

// 	merge(nums1, m, nums2, n)
// }
