package main

import (
	"strings"
)

func licenseKeyFormatting(s string, k int) string {
	var (
		temp, temp1, res string
		count            int
	)

	for _, v := range s {
		if string(v) == "-" {
			continue
		}
		temp += strings.ToUpper(string(v))
	}
	for i := len(temp) - 1; i >= 0; i-- {
		temp1 += string(temp[i])
		count++
		if count == k && i != 0 {
			temp1 += "-"
			count = 0
		}
	}
	for i := len(temp1) - 1; i >= 0; i-- {
		res += string(temp1[i])
	}
	return res
}

// func main() {
// 	s := "2-5g-3-J"
// 	k := 2

// 	fmt.Println(licenseKeyFormatting(s, k))
// }
