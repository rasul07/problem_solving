package main

func findDisappearedNumbers(nums []int) []int {
	var res []int
	// carry := 1
	m := make(map[int]int)

	for _, v := range nums {
		m[v]++
	}
	// for i := 0; i < len(nums); i++ {
	// 	if nums[i] != carry {
	// 		res = append(res, i)
	// 		continue
	// 	}
	// 	carry++
	// }

	return res
}

// func main() {
// 	nums := []int{1, 1}

// 	fmt.Println(findDisappearedNumbers(nums))
// }
