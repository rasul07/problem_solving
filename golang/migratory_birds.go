package main

func migratoryBirds(arr []int32) int32 {
	// Write your code here
	var (
		birdType int32
		m        = map[int32]int32{}
		keys     []int32
	)
	for _, v := range arr {
		m[v]++
	}
	max := m[0]
	for _, v := range m {
		if v > max {
			max = v
		}
	}

	for k, v := range m {
		if v >= max {
			max = v
			keys = append(keys, k)
		}
	}
	birdType = keys[0]
	for _, v := range keys {
		if v < birdType {
			birdType = v
		}
	}
	return birdType
}

// func main() {
// 	arr := []int32{1, 4, 4, 4, 5, 3}

// 	fmt.Println(migratoryBirds(arr))
// }
