package main

import "fmt"

func miniMaxSum(arr []int32) {
	// Write your code here
	var (
		mi, ma        int64 = int64(arr[0]), int64(arr[0])
		min, max, sum int64
	)
	for _, v := range arr {
		if int64(v) < mi {
			mi = int64(v)
		}
		if int64(v) > ma {
			ma = int64(v)
		}
	}
	fmt.Println("Mi ==> ", mi)
	fmt.Println("Ma ==> ", ma)
	for _, v := range arr {
		sum += int64(v)
	}

	fmt.Println("Sum ==> ", sum)

	min = sum - mi
	max = sum - ma
	fmt.Println("Min ==> ", min)
	fmt.Println("Max ==> ", max)
	fmt.Print(max, " ", min)
}

func main() {
	arr := []int32{156873294, 719583602, 581240736, 605827319, 895647130}

	miniMaxSum(arr)
}
