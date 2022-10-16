package main

func containsDuplicate(nums []int) bool {
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
		if m[nums[i]] == 2 {
			return true
		}
	}
	return false
}

// func main() {
// 	nums := []int{1,1,1,3,3,4,3,2,4,2}

// 	fmt.Println(containsDuplicate(nums))
// }
