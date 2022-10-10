package main

import (
	"fmt"
	"time"
)

func times() {

	time.Sleep(300 * time.Millisecond)
}
func times2() {
	time.Sleep(400 * time.Millisecond)
	for i := 0; i < 100; i++ {
		for j := 0; j < i; j++ {
			fmt.Println("Nothing")
		}
	}
}
func times3() {
	time.Sleep(500 * time.Millisecond)
}
// func main() {
// 	type str struct {
// 		name string
// 		age  int32
// 	}

// 	m := make(map[float64]str)

// 	m[4.5] = str{
// 		"Rasul",
// 		19,
// 	}
// 	n := time.Now()
// 	go times()
// 	go times2()
// 	go times3()
// 	fmt.Printf("execution time : %s", time.Since(n))
// }