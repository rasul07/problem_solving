package main

func intersection(nums1 []int, nums2 []int) []int {
	var res []int
	m := make(map[int]int)
	m1 := make(map[int]int)

	for _, v := range nums1 {
		if _, ok := m[v]; !ok {
			m[v]++
		}
	}
	for _, v := range nums2 {
		if _, ok := m[v]; ok {
			m1[v]++
		}
	}

	for k := range m1 {
		res = append(res, k)
	}

	return res
}

// func main() {
// 	nums1 := []int{1}
// 	nums2 := []int{1}

// 	fmt.Println(intersection(nums1, nums2))
// }
