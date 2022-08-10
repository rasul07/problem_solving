package main

func birthdayCakeCandles(candles []int32) int32 {
	// Write your code here
	var (
		count int32
		m     = map[int32]int32{}
		x     = candles[0]
	)

	for _, v := range candles {
		m[v]++
	}
	// fmt.Println(m)
	for i := range m {
		// fmt.Println("Here ", i, " ", v)
		if i > x {
			x = i
		}
	}
	// fmt.Println(x)
	count = m[x]
	return count
}

// func main() {
// 	candles := []int32{4, 2, 3, 4, 5}

// 	fmt.Println(birthdayCakeCandles(candles))
// }
