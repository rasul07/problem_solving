package main

import (
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var median float64

	nums1 = append(nums1, nums2...)

	sort.Ints(nums1)
	l := len(nums1)

	if l%2 == 0 {
		median = float64((float64(nums1[l/2-1]) + float64(nums1[l/2])) / 2)
	} else {
		median = float64(nums1[l/2])
	}

	return median
}

// func main() {
// 	n1 := []int{1, 2}
// 	n2 := []int{3, 4}

// 	fmt.Println(findMedianSortedArrays(n1, n2))
// }
