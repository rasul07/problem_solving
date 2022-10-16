package main

func generate(numRows int) [][]int {
	var (
		res = make([][]int, numRows)
	)
	if numRows == 1 {
		res[0] = []int{1}
		return res
	} else if numRows == 2 {
		res[0] = []int{1}
		res[1] = []int{1, 1}
		return res
	}

	res[0] = []int{1}
	res[1] = []int{1, 1}

	for i := 2; i < numRows; i++ {
		temp := make([]int, i+1)
		temp[0] = 1
		temp[i] = 1
		for j := 1; j < i; j++ {
			temp[j] = res[i-1][j-1] + res[i-1][j]
		}
		res[i] = temp
	}
	return res
}

// func main() {
// 	n := 10
// 	fmt.Println(generate(n))
// }
