package main

func largestAltitude(gain []int) int {
	var (
		pref, max int
	)

	for _, v := range gain {
		pref += v
		if pref > max {
			max = pref
		}
	}

	return max
}

// func main() {
// 	arr := []int{-4,-3,-2,-1,4,3,2}

// 	fmt.Println(largestAltitude(arr))
// }
