package main

import (
	"strings"
)

func timeConversion(s string) string {
	// Write your code here
	l := len(s)
	lastDigits := s[(l - 2):]
	firstDigits := s[0:2]
	result := ""
	m := map[string]string{
		"01": "13",
		"02": "14",
		"03": "15",
		"04": "16",
		"05": "17",
		"06": "18",
		"07": "19",
		"08": "20",
		"09": "21",
		"10": "22",
		"11": "23",
	}
	if lastDigits == "PM" && firstDigits != "12" {
		result = s[0 : l-2]
		result = strings.Replace(result, firstDigits, m[firstDigits], 1)
	} else if lastDigits == "AM" && firstDigits == "12" {
		result = s[0 : l-2]
		result = strings.Replace(result, firstDigits, "00", 1)
	} else {
		result = s[0 : l-2]
	}
	return result
}

// func main() {
// 	fmt.Println(timeConversion("12:45:54PM"))
// }
