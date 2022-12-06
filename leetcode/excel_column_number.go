package main

import (
	"math"
)

func titleToNumber(columnTitle string) int {
	var res int
	m := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"D": 4,
		"E": 5,
		"F": 6,
		"G": 7,
		"H": 8,
		"I": 9,
		"J": 10,
		"K": 11,
		"L": 12,
		"M": 13,
		"N": 14,
		"O": 15,
		"P": 16,
		"Q": 17,
		"R": 18,
		"S": 19,
		"T": 20,
		"U": 21,
		"V": 22,
		"W": 23,
		"X": 24,
		"Y": 25,
		"Z": 26,
	}
	l := len(columnTitle) - 1
	if len(columnTitle) > 1 {
		for i := 0; i < l; i++ {
			res += int(math.Pow(float64(m["Z"]), float64(l-i))) * m[string(columnTitle[i])]
		}
		res += m[string(columnTitle[l])]
	} else {
		res = m[columnTitle]
	}

	return res
}

// func titleToNumber(columnTitle string) int {
//     res := 0
//     for i, _ := range columnTitle {
//         res *= 26
//         res += int(columnTitle[i] - 'A') + 1
//     }
//     return res
// }

// func main() {
// 	c := ""
// 	fmt.Println(2147483647)
// 	fmt.Println(titleToNumber(c))
// }
