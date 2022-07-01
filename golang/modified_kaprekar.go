package main

import (
	"fmt"
	"math"
)

func digits(n int64) int64 {
	var (
		res int64
		i   = n
	)
	for i > 0 {
		i /= 10
		res++
	}
	return res
}

func isKaprekar(x int64) bool {
	// fmt.Println("XXXXX =====> ", x)
	var (
		l, r, sum, d int64
	)
	d = digits(x)
	// fmt.Println("X ===>", x,"DIGITS =====> ", d)
	l = int64(x*x) / int64(math.Pow(10, float64(d)))
	// fmt.Println("X2 =====> ", int64(x*x))
	// fmt.Println("Left ===> ", l)
	r = int64(x*x) % int64(math.Pow(10, float64(d)))
	// fmt.Println("Right ===> ", r)
	sum = l + r
	// fmt.Println("SUM ====>", sum)
	if sum == x {
		return true
	} else {
		return false
	}
}

func kaprekarNumbers(p int32, q int32) {
	// Write your code here
	var (
		nums    []int32
		invalid string
	)
	for i := p; i <= q; i++ {
		if isKaprekar(int64(i)) {
			// fmt.Println("TRUE =====>", i)
			nums = append(nums, i)
		} else {
			invalid = "INVALID RANGE"
		}
	}

	if len(nums) != 0 {
		for _, v := range nums {
			fmt.Print(v, " ")
		}
	} else {
		fmt.Println(invalid)
	}

}

func main() {
	// kaprekarNumbers(1, 99999)
	n := 99999 * 99999

	// fmt.Println(n, digits(int32(n)))
	left := uint64(n) / uint64(math.Pow(10, float64(5)))
	right := uint64(n) % uint64(math.Pow(10, float64(5)))

	fmt.Println(n, "Digit +=====> ", digits(int64(n)), left, right, left+right)
	fmt.Println(isKaprekar(99999))
}
