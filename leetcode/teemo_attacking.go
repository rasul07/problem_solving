package main

func findPoisonedDuration(timeSeries []int, duration int) int {
	var total int

	for i := 0; i < len(timeSeries)-1; i++ {
		if timeSeries[i+1]-timeSeries[i] < duration {
			total += timeSeries[i+1] - timeSeries[i]
		} else {
			total += duration
		}
	}

	return total + duration
}

// func main() {
// 	t := []int{1, 4}
// 	d := 2

// 	fmt.Println(findPoisonedDuration(t, d))
// }
