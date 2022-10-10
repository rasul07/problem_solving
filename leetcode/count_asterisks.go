package main

func countAsterisks(s string) int {
	var (
		result, count int
	)
	for _, v := range s {
		if string(v) == "|" {
			count++
		}
		if count == 3 {
			count = 1
		}
		if count == 2 || count == 0 {
			if string(v) == "*" {
				result++
			}
		}
	}
	return result
}

// func main() {
// 	s := "*||*|||||*|*|***||*||***|"

// 	fmt.Println(countAsterisks(s))

// }
