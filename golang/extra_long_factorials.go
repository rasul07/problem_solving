package main

import (
	"fmt"
	"math/big"
)

func extraLongFactorials(n int32) {
	// Write your code here
	// var sum int32 = n
	// for i := 1; i < int(n); i++ {
	// 	sum *= (n - int32(i))
	// }
	// fmt.Println(uint(sum))
	fact := big.NewInt(1)

	for i := 1; i <= int(n); i++ {
		fact.Mul(fact, big.NewInt(int64(i)))
	}

	fmt.Println(fact)
}

func main() {
	for i := 1; i <= 25; i++ {
		fmt.Printf("%v ", i)
		extraLongFactorials(int32(i))
	}
}
