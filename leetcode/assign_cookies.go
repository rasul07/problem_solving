package main

// func quickSort(arr []int) []int {
// 	if len(arr) < 2 {
// 		return arr
// 	}

// 	left, right := 0, len(arr)-1
// 	pivot := (left + right) / 2

// 	arr[pivot], arr[right] = arr[right], arr[pivot]

// 	for i := range arr {
// 		if arr[i] < arr[right] {
// 			arr[left], arr[i] = arr[i], arr[left]
// 			left++
// 		}
// 	}

// 	arr[left], arr[right] = arr[right], arr[left]
// 	quickSort(arr[:left])
// 	quickSort(arr[left+1:])

// 	return arr
// }

func findContentChildren(g []int, s []int) int {
	var (
		res int
	)
	j := 0
	i := 0
	quickSort(g)
	quickSort(s)

	for i < len(s) && j < len(g) {
		if s[i] >= g[j] {
			res++
			j++
		}
		i++
	}

	return res
}

// func main() {
// 	g := []int{250, 490, 328, 149, 495, 325, 314, 360, 333, 418, 430, 458}
// 	s := []int{1234123}
// 	fmt.Println(findContentChildren(g, s))
// }
