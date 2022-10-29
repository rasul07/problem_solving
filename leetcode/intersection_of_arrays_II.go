package main

func intersect(nums1 []int, nums2 []int) []int {
	var (
		res []int
	)
	m := make(map[int]int)
	m1 := make(map[int]int)

	for _, v := range nums1 {
		m[v]++
	}
	for _, v := range nums2 {
		if _, ok := m[v]; ok {
			m1[v]++
			if m[v] < m1[v] {
				m1[v] = m[v]
			}
		}
	}

	for k, v := range m1 {
		for v > 0 {
			res = append(res, k)
			v--
		}
	}

	return res
}

// func main() {
// 	nums1 := []int{2, 2}
// 	nums2 := []int{1, 2, 2, 2, 1}

// 	fmt.Println(intersect(nums1, nums2))
// }
