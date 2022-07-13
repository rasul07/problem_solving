package main

import "fmt"

func timeInWords(h int32, m int32) string {
    // Write your code here
	var result string
	min, mins, past, to := "minute", "minutes", "past", "to"
	
	strs := map[int32]string{
		0: "o' clock",
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
		6: "six",
		7: "seven",
		8: "eight",
		9: "nine",
		10: "ten",
		11: "eleven",
		12: "twelve",
		13: "thirteen",
		14: "fourteen",
		15: "quarter",
		16: "sixteen",
		17: "seventeen",
		18: "eighteen",
		19: "nineteen",
		20: "twenty",
		21:	"twenty one",
		22:	"twenty two",
		23:	"twenty three",
		24:	"twenty four",
		25:	"twenty five",
		26:	"twenty six",
		27:	"twenty seven",
		28:	"twenty eight",	
		29:	"twenty nine",
		30: "half",
	}

	if m == 0 {
		result = strs[h] + " " + strs[0]
		return result
	} else {
		if m <= 30 {
			if m == 1 {
				result = strs[m] + " " + min + " " + past + " " + strs[h]
			} else if m == 15 || m == 30 {
				result = strs[m] + " " + past + " " + strs[h]
			} else {
				result = strs[m] + " " + mins + " " + past + " " + strs[h]
			}	
		} else if m > 30 {
			if 60 - m == 1 {
				result = strs[60 - m] + " " + min + " " + to + " " + strs[h + 1]
			} else if 60 - m == 15 {
				result = strs[60 - m] + " " + to + " " + strs[h + 1]
			} else {
				result = strs[60-m] + " " + mins + " " + to + " " + strs[h + 1]
			}
		}
		return result
	}

}

func main() {
	var (
		h int32 = 5
		m int32 = 30
	) 

	fmt.Println(timeInWords(h, m))
}