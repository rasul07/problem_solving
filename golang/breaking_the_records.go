package main

func breakingRecords(scores []int32) []int32 {
	// Write your code here
	var (
		recMaxs, recMins = scores[0], scores[0]
		maxs, mins       int32
	)
	results := []int32{}

	for _, v := range scores {
		if v > recMaxs {
			recMaxs = v
			maxs++
		}
		if v < recMins {
			recMins = v
			mins++
		}
	}

	results = append(results, maxs, mins)
	return results
}

// func main() {
// 	arr := []int32{12, 24, 10, 24}

// 	fmt.Println(breakingRecords(arr))
// }
