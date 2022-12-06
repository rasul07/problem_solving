package main

func lengthOfLastWord(s string) int {
	var (
		str    []string
		temp   string
		length int
	)
	if len(s) == 1 {
		return 1
	}
	for i := 1; i < len(s); i++ {
		if i == len(s)-1 {
			if string(s[i]) != " " && string(s[i-1]) != " " {
				temp += string(s[i-1])
				temp += string(s[i])
				str = append(str, temp)
			} else if string(s[i]) != " " && string(s[i-1]) == " " {
				temp += string(s[i])
				str = append(str, temp)
			}
		}
		if string(s[i]) != " " && string(s[i-1]) != " " {
			temp += string(s[i-1])
		} else if string(s[i]) == " " && string(s[i-1]) != " " {
			temp += string(s[i-1])
			str = append(str, temp)
			temp = ""
		}
	}

	length = len(str[len(str)-1])

	return length
}

// func main() {
// 	s := " H"

// 	fmt.Println(lengthOfLastWord(s))
// }
