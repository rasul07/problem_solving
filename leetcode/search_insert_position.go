package main

func searchInsert(nums []int, target int) int {
	var resp int
	for i := 0; i < len(nums); i++ {
		if target > nums[len(nums)-1] {
			resp = len(nums)
			break
		} else if target < nums[i] {
			resp = i
			break
		} else if target == nums[i] {
			resp = i
			break
		}
	}
	return resp
}

// func main() {
// 	nums := []int{1, 3, 5, 6}
// 	target := 1

// 	fmt.Println(searchInsert(nums, target))
// }
