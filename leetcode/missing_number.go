package main

func missingNumber(nums []int) int {
	var res int
	m := make(map[int]int)

	for _, v := range nums {
		m[v]++
	}
	for i := 0; i <= len(nums); i++ {
		if _, ok := m[i]; !ok {
			res = i
		}
	}

	return res
}

// func main() {
// 	nums := []int{1, 2}

// 	fmt.Println(missingNumber(nums))
// }
