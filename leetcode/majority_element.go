package main

func majorityElement(nums []int) int {
	var res int
	m := make(map[int]int)
	l := len(nums)
	for _, v := range nums {
		m[v]++
		if m[v] > l/2 {
			res = v
			break
		}
	}
	return res
}

// func main() {
// 	nums := []int{3,2,3}

// 	fmt.Println(majorityElement(nums))
// }
