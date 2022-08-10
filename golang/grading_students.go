package main

func gradingStudents(grades []int32) []int32 {
	// Write your code here
	var (
		finals []int32
		next   int32
	)

	for _, v := range grades {
		next = 5 - v%5
		if v >= 38 {
			if next < 3 {
				finals = append(finals, next+v)
			} else {
				finals = append(finals, v)
			}
		} else {
			finals = append(finals, v)
		}
	}
	return finals
}

// func main() {
// 	g := []int32{73, 67, 38, 33}

// 	fmt.Println(gradingStudents(g))
// }
