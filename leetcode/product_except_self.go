package main

func productExceptSelf(nums []int) []int {
	//make array with same length as nums
	res := make([]int, len(nums))
	//create variable for computing product
	prod := 1
	//set first element of res to 1
	res[0] = 1
	//first loop will compute products of elements before each element
	for i := 1; i < len(nums); i++ {
		prod *= nums[i-1]

		res[i] = prod
	}
	//reset product
	prod = 1
	//second loop will compute products of elements after each element
	for i := len(nums) - 2; i >= 0; i-- {
		prod *= nums[i+1]
		res[i] *= prod
	}
	return res
}

// func main() {
// 	n := []int{5, 7, 3, 4 , 2}

// 	fmt.Println(productExceptSelf(n))
// }
