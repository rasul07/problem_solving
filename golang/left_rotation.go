package main

func rotateLeft(d int32, arr []int32) (newArr []int32) {
	// Write your code here

	if d == 0 {
		return arr
	}
	if int(d) > len(arr) {
		return arr
	}
	var aArr = arr[:d]
	var bArr = arr[d:cap(arr)]
	arr = nil
	// fmt.Println(arr)
	arr = append(arr, bArr...)
	arr = append(arr, aArr...)

	for _, val := range arr {
		if val != 0 {
			newArr = append(newArr, val)
		}
	}

	return newArr
}

// func main() {
// 	var d int32 = 4
// 	arr := []int32{1, 2, 3, 4, 5}

// 	fmt.Println(rotateLeft(d, arr))
// }
