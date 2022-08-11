package main

import "fmt"

func bonAppetit(bill []int32, k int32, b int32) {
	// Write your code here
	var sum int32
	for _, v := range bill {
		if v == bill[k] {
			continue
		}
		sum += v
	}
	if sum/2 == b {
		fmt.Println("Bon Appetit")
	} else {
		fmt.Println(b-sum/2)
	}
}

// func main() {
// 	bill := []int32{3, 10, 2, 9}
// 	k := int32(1)
// 	b := int32(12)

// 	bonAppetit(bill, k, b)
// }
