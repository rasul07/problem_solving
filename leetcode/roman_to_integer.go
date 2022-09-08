package main

func romanToInt(s string) (res int) {
	m := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	for i := 0; i < len(s); i++ {
		if i != len(s)-1 {
			if m[string(s[i])] < m[string(s[i+1])] {
				res += m[string(s[i+1])] - m[string(s[i])]
				i++
			} else {
				res += m[string(s[i])]
			}
		} else {
			res += m[string(s[i])]
		}
	}
	return
}

// func main() {
// 	fmt.Println(romanToInt("XV"))

// }
