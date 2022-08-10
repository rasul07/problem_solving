package main

func countArray(a []int32) map[int32]int32 {
	c := make(map[int32]int32)
	for _, i := range a {
		c[i] += 1
	}

	return c
}

func maxInt32(a int32, b int32) int32 {
	if a > b {
		return a
	}

	return b
}

func pickingNumbers(a []int32) int32 {
	var pn int32 = 0
	c := countArray(a)

	for k, v := range c {
		pn = maxInt32(pn, v+c[k+1])
	}

	return pn
}

// func main() {

// 	a := []int32{1, 1, 2, 2, 4, 4, 5, 5, 5}
// 	fmt.Println(pickingNumbers(a))
// }
