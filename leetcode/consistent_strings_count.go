package main

import (
	"strings"
)

// NOT DONE
func countConsistentStrings(allowed string, words []string) int {
	var (
		result int
	)

	for i := 0; i < len(words); i++ {
		arr := strings.Split(words[i], "")

		for j := 0; j < len(arr); j++ {
			if string(allowed[j]) != arr[j] {
				break
			}
		}
		result++
	}

	return result
}

// func main() {
// 	var (
// 		allowed string = "cad"
// 		words          = []string{"cc","acd","b","ba","bac","bad","ac","d"}
// 	)

// 	fmt.Println(countConsistentStrings(allowed, words))
// }
