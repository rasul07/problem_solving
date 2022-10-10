package main

func removeElement(nums []int, val int) int {
	res := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[res] = nums[i]
			res++
		}
	}

	return res
}

// func main() {
// 	arr := []int{0, 1, 2, 2, 3, 0, 4, 2}

// 	fmt.Println(removeElement(arr, 2))
// }
