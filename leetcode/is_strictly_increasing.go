package main
// NOT DONE
func removedArr(nums []int, x int) (res []int) {
	for _, v := range nums {
		if v == x {
			continue
		}
		res = append(res, v)
	}

	return
}

func isIncreasing(nums []int) bool {
	var (
		result bool = true
		l      int  = len(nums) - 1
	)

	for i := l - 1; i > 0; i-- {
		if nums[i+1] > nums[i] && nums[i] > nums[i-1] {
			continue
		} else {
			result = false
			break
		}
	}

	return result
}

func canBeIncreasing(nums []int) bool {
	var result bool

	for i := 0; i < len(nums); i++ {
		result = isIncreasing(removedArr(nums, nums[i]))
	}

	return result
}

// func main() {
// 	nums := []int{1, 2, 10, 5, 7}
// 	// nums := []int{1, 1, 1}

// 	fmt.Println(canBeIncreasing(nums))
// }
