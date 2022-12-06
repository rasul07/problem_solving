package main

func twoSum(nums []int, target int) (resp []int) {
	// for i := 0; i < len(nums); i++ {
	// 	for j := i + 1; j < len(nums); j++ {
	// 		if nums[i] + nums[j] == target {
	// 			resp = append(resp, i, j)
	// 			break
	// 		}
	// 	}
	// 	if len(resp) == 2 {
	// 		break
	// 	}
	// }
	m := map[int]int{}
	for i, v := range nums {
		if _, ok := m[target-v]; ok {
			resp = append(resp, m[target-v], i)
			return resp
		}
		m[v] = i
	}
	return resp
}

// func main() {
// 	var (
// 		nums   = []int{3, 3}
// 		target = 6
// 	)

// 	fmt.Println(twoSum(nums, target))
// }
