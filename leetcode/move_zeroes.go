package main

func moveZeroes(nums []int) []int {
	l := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[l] = nums[i]
			nums[i] = 0
			l++
		}
	}

	return nums
}

// func main() {
// 	arr := []int{0, 1, 0, 3, 12}

// 	fmt.Println(moveZeroes(arr))
// }
