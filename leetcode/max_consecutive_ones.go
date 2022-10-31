package main

func findMaxConsecutiveOnes(nums []int) int {
	var max, ones int

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			ones++
		} else {
			if ones > max {
				max = ones
			}
			ones = 0
		}
	}
	if ones > max {
		max = ones
	}

	return max
}

// func main() {
// 	n := []int{1,1, 0, 1, 1, 1, 1}
// 	fmt.Println(findMaxConsecutiveOnes(n))
// }
