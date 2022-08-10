package main

func jumpingOnClouds(c []int32, k int32) int32 {
	var (
		e int32 = 100
	)
	for i := 0; i < len(c); i += int(k) {
		if c[i] == 0 {
			e--
		} else {
			e -= 3
		}

	}
	return e
}

// func main() {
// 	arr := []int32{0, 1, 1, 0, 1, 0, 1, 1, 0, 1, 1, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 0, 1, 0}
// 	var k int32 = 3

// 	fmt.Println(jumpingOnClouds(arr, k))
// }
