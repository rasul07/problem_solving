package main

import "fmt"

func dayOfProgrammer(year int32) string {
	// Write your code here
	var result string
	if year == 1918 {
		result = "28.08.1918"
	} else if (year%4 == 0 && year%100 != 0) || year%400 == 0 || (year < 1918 && year%4 == 0) {
		result = fmt.Sprintf("12.09.%v", year)
	} else {
		result = fmt.Sprintf("13.09.%v", year)
	}

	return result
}

// func main() {
// 	y := 1800
// 	fmt.Println(dayOfProgrammer(int32(y)))
// }
