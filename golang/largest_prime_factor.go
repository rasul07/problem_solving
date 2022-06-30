package main

import "fmt"

func main() {
	n := int64(600851475143)
	var p, p1, j int64
	for i := 0; int64(i) < n/2; i++ {
		p = 6*int64(i) - 1
		if n%p == 0 {
			j = p
		}
		p1 = 6*int64(i) + 1
		if n%p == 0 {
			j = p1
		}
	}
	fmt.Println(j)
}