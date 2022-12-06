package main

import (
	"strconv"
	"strings"
)

func squareIsWhite(coordinates string) bool {
	var (
		a = []bool{false, true, false, true, false, true, false, true}
		b = []bool{true, false, true, false, true, false, true, false}
	)
	arr := strings.Split(coordinates, "")
	ind, _ := strconv.Atoi(arr[1])

	switch arr[0] {
	case "a", "c", "e", "g":
		return a[ind-1]
	case "b", "d", "f", "h":
		return b[ind-1]
	}
	return true
}

// func main() {
// 	s := "h8"

// 	fmt.Println(squareIsWhite(s))
// }
