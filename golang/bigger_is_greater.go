package main

import (
	"fmt"
	"strings"
)

func noAnswer(str string) bool {
	l, r := len(str)-1, false
	i := 1
	fmt.Println(str, i)
	if l - i > -1 {
		noAnswer(strings.Trim(str, string(str[l])))
		if str[l] <= str[l-i] {
			r = true
		} else {
			r = false
		}
		i++
	}
	
	return r
}

func biggerIsGreater(w string) string {
	// Write your code here
	result := ""

	return result
}

func main() {
	w := "cdcca"
	// fmt.Println(biggerIsGreater(w))
	// fmt.Println(noAnswer(w))
	fmt.Println(noAnswer(w))
}
