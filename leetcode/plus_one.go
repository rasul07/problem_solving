package main

func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return digits
	}
	carry := 1
	for j := len(digits) - 1; j >= 0; j-- {
		sum := digits[j] + carry
		digits[j] = sum % 10
		carry = sum / 10
	}
	if carry == 1 {
		digits = append([]int{1}, digits...)
	}
	return digits
}

// func main() {
// 	arr := []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 9, 9}

// 	fmt.Println(plusOne(arr))
// }
