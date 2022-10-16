package main

func findSubarrays(nums []int) bool {
	m := make(map[int]int)

	for i := 1; i < len(nums); i++ {
		sum := nums[i-1] + nums[i]
		m[sum]++
		if m[sum] > 1 {
			return true
		}
	}
	return false
}

// func main() {
// 	nums := []int{0,0}

// 	fmt.Println(findSubarrays(nums))
// }
