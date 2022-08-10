package main

import "fmt"

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	// Write your code here
	var apple, orange int
	for _, v := range apples {
		if s <= a+v && a+v <= t {
			apple++
		}
	}
	for _, v := range oranges {
		if s <= b+v && t >= b+v  {
			fmt.Println(b+v)
			orange++
		}
	}
	fmt.Println(apple)
	fmt.Println(orange)
}

// func main() {
// 	var (
// 		s, t, a, b      int32   = 2, 3, 1, 5
// 		apples, oranges []int32 = []int32{-2}, []int32{-1}
// 	)
// 	countApplesAndOranges(s, t, a, b, apples, oranges)
// }
