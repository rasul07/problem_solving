package main

func singleNumber(nums []int) int {
	m := make(map[int]int)
	var res int
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	for _, v := range nums {
		if m[v] == 1 {
			res = v
		}
	}
	return res
}

// func main() {
// 	nums := []int{1}

// 	fmt.Println(singleNumber(nums))
// }
