package main

import "fmt"

func plusMinus(arr []int32) {
	// Write your code here
	var (
		pos  float32
		neg  float32
		zero float32
		l    float32 = float32(len(arr))
	)
	for _, v := range arr {
		if v > 0 {
			pos++
		} else if v < 0 {
			neg++
		} else if v == 0 {
			zero++
		}
	}
	pos /= l
	neg /= l
	zero /= l

	fmt.Printf("%.6f\n", pos)
	fmt.Printf("%.6f\n", neg)
	fmt.Printf("%.6f\n", zero)

}

func main() {
	var arr = []int32{1, 1, 0, -1, -1}

	plusMinus(arr)
}
