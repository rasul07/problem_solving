package main

func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	// Write your code here
	var pairs int32
	for i := 0; i < len(ar); i++ {
		for j := i + 1; j < len(ar); j++ {
			if (ar[i]+ar[j])%k == 0 {
				pairs++
			}
		}

	}
	return pairs
}

// func main() {
// 	n := int32(6)
// 	k := int32(5)
// 	arr := []int32{1, 2, 3, 4, 5, 6}

// 	fmt.Println(divisibleSumPairs(n, k, arr))
// }
