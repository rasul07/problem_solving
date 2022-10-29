package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	var (
		patArr, sArr []string
		m = make(map[string]string)
		m1 = make(map[string]string)
	)

	patArr = strings.Split(pattern, "")
	sArr = strings.Split(s, " ")
	m[patArr[0]] = sArr[0]
	m1[sArr[0]] = patArr[0]
	fmt.Println(m, m1)
	for i := 1; i < len(patArr); i++ {
		if _, ok := m[patArr[i]]; !ok {
			if v, ok := m1[sArr[i]]; ok && v != patArr[i] {
				return false
			} else {
				m[patArr[i]] = sArr[i]
				m1[sArr[i]] = patArr[i]
			}
		} else if m1[sArr[i]] != patArr[i] {
			return false
		}
	}
	return true
}

func main() {
	pattern := "aaaa"
	s := "dog dog dog dog"

	fmt.Println(wordPattern(pattern, s))
}
