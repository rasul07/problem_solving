package main

import (
	"fmt"
)

func lenInt(i int32) []int32 {
	var res []int32

	for i > 0 {
		res = append(res, i%10)
		i /= 10
	}

	return res
}

func findDigits(n int32) int32 {
	// Write your code here
	var res int32

	l := lenInt(n)

	for _, v := range l {
		if v == 0 {
			continue
		} else if n%v == 0 {
			res++
		}
	}

	return res
}

func main() {
	t := 1012

	fmt.Println(findDigits(int32(t)))
}
