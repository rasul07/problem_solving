package main

func lenInt(i int32) []int32 {
	var res []int32

	for i > 0 {
		res = append(res, i%10)
		i /= 10
	}

	return res
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	arr := lenInt(int32(x))
	l := len(arr)

	for i := 0; i < l/2; i++ {
		if arr[i] != arr[l-1-i] {
			return false
		}
	}

	return true
}

// func main() {
// 	fmt.Println(isPalindrome(10))
// }
