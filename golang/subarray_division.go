package main

func sumOfElements(x []int32) int32 {
	var result int32
	for _, v := range x {
		result += v
	}
	return result
}

func birthday(s []int32, d int32, m int32) int32 {
	// Write your code here
	var possibleWays int32

	for i := 0; i < len(s); i++ {
		if (sumOfElements(s[i:(int32(i) + m)])) == d {
			possibleWays++
		}
		if int32(i)+m >= int32(len(s)) {
			break
		}
	}

	return possibleWays
}

// func main() {
// 	var (
// 		s = []int32{1, 1, 1, 1, 1}
// 		d = int32(4)
// 		m = int32(2)
// 	)

// 	fmt.Println(birthday(s, d, m))
// }
