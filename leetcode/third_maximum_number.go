package main

func thirdMax(nums []int) int {
	var (
		firstMax              int = nums[0]
		secondMax, thirdMax   int
		maxCount, secMaxCount int
	)
	for _, v := range nums {
		if v < thirdMax {
			thirdMax = v
			secondMax = v
		}
	}
	for _, v := range nums {
		if v > firstMax {
			firstMax = v
		}
	}
	for _, v := range nums {
		if v < firstMax && v > secondMax {
			secondMax = v
			secMaxCount++
		}
	}
	for _, v := range nums {
		if v < secondMax && v >= thirdMax && secMaxCount != 0 {
			thirdMax = v
			maxCount++
		}
	}

	if maxCount >= 1 {
		return thirdMax
	}
	return firstMax
}

// func main() {
// 	n := []int{3,2,1}

// 	fmt.Println(thirdMax(n))
// }
