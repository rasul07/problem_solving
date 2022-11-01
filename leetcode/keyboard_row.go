package main

func findWords(words []string) []string {
	var (
		c1, c2, c3 int
		res        []string
	)
	row1 := map[string]bool{
		"Q": true,
		"W": true,
		"E": true,
		"R": true,
		"T": true,
		"Y": true,
		"U": true,
		"I": true,
		"O": true,
		"P": true,
		"q": true,
		"w": true,
		"e": true,
		"r": true,
		"t": true,
		"y": true,
		"u": true,
		"i": true,
		"o": true,
		"p": true,
	}
	row2 := map[string]bool{
		"A": true,
		"S": true,
		"D": true,
		"F": true,
		"G": true,
		"H": true,
		"J": true,
		"K": true,
		"L": true,
		"a": true,
		"s": true,
		"d": true,
		"f": true,
		"g": true,
		"h": true,
		"j": true,
		"k": true,
		"l": true,
	}
	row3 := map[string]bool{
		"Z": true,
		"X": true,
		"C": true,
		"V": true,
		"B": true,
		"N": true,
		"M": true,
		"z": true,
		"x": true,
		"c": true,
		"v": true,
		"b": true,
		"n": true,
		"m": true,
	}

	for _, v := range words {
		c1, c2, c3 = 0, 0, 0
		for _, l := range v {
			if _, ok := row1[string(l)]; ok {
				c1++
				if c1 == len(v) {
					res = append(res, v)
					break
				}
			} else if _, ok := row2[string(l)]; ok {
				c2++
				if c2 == len(v) {
					res = append(res, v)
					break
				}
			} else if _, ok := row3[string(l)]; ok {
				c3++
				if c3 == len(v) {
					res = append(res, v)
					break
				}
			}
		}
	}
	return res
}

// func main() {
// 	w := []string{"qwer", "asdq", "zxcvq", "vcvb"}
// 	fmt.Println(findWords(w))
// }
