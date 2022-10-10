package main

func removeDuplicates(nums []int) int {
	pos := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[pos] {
			pos++
			tmp := nums[i]
			nums[i] = nums[pos]
			nums[pos] = tmp
		}
	}
	return pos + 1
}

// func main() {
// 	arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

// 	fmt.Println(removeDuplicates(arr))
// }
