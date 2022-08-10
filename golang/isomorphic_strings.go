package main

func stringToMap(s string) string {
	var (
		myMap      = make(map[rune]rune)
		a     rune = 65
		str   string
	)

	for _, v := range s {
		if _, ok := myMap[v]; ok {
			str += string(myMap[v])
		} else {
			myMap[v] = a
			str += string(myMap[v])
			a++
		}
	}
	return str
}

func isIsomorphic(s string, t string) bool {
	return stringToMap(s) == stringToMap(t)
}
