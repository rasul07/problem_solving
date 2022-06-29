package main

import "fmt"

func hourglassSum(arr [][]int32) int32 {
    // Write your code here
	var sum1 int32
	
	sum := arr[0][0] + arr[0][1] + arr[0][2] + arr[1][1] + arr[2][0] + arr[2][1] + arr[2][2]
	for i, v := range arr {
		// if i == 1 {
		// 	break
		// }
		for j, k := range v {
			if j + 2 < len(v) && i+2 < len(arr) {
				sum1 = k + v[j+1] + v[j+2] + arr[i+1][j+1] + arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]
				// fmt.Println("First", k)
				// fmt.Println("Second", v[j+1])
				// fmt.Println("Third", v[j+2])
				// fmt.Println("Fourth", arr[i+1][j+1])
				// fmt.Println("Fifth", arr[i+2][j])
				// fmt.Println("Sixth", arr[i+2][j+1])
				// fmt.Println("Seventh", arr[i+2][j+2])
			    // fmt.Println("Sum ======>>", sum1)
			}
			if sum1 > sum {
				sum = sum1
			}
		}
	}
	return sum
}

func main() {
	var arr [][]int32 = [][]int32{
		{-1, -1, 0, -9, -2, -2},
		{-2, -1, -6, -8, -2, -5},
		{-1, -1, -1, -2, -3, -4},
		{-1, -9, -2, -4, -4, -5},
		{-7, -3, -3, -2, -9, -9},
		{-1, -3, -1, -2, -4, -5},
	}

	fmt.Println("main", hourglassSum(arr))
}