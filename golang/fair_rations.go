package main

func fairRations(B []int32) string {
	// Write your code here
	var count int32

	for i := 0; i < len(B); i++ {
		if i == 0 && B[i]%2 != 0 {
			count++
		} else if i == len(B) - 1 && B[i]%2 != 0 {
			count++
		}
	}
	return "Yes"
}

func main() {
	b := []int32{2, 3, 4, 5, 6}

	fairRations(b)
}
