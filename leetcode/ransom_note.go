package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	note := make(map[string]int)
	mgzn := make(map[string]int)

	for _, v := range ransomNote {
		note[string(v)]++
	}

	for _, v := range magazine {
		mgzn[string(v)]++
	}

	for k, v := range note {
		if _, ok := mgzn[k]; ok {
			if mgzn[k] >= v {
				fmt.Println(mgzn[k], v)
				continue
			} else {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

// func main() {
// 	r := "a"
// 	m := "b"

// 	fmt.Println(canConstruct(r, m))
// }
