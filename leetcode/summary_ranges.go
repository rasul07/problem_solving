package main

import (
	"strconv"
)

func summaryRanges(nums []int) []string {
	var (
		result        []string
		str           string
		countSequence int
	)

	for i := 0; i < len(nums); i++ {

		if i < len(nums)-1 && nums[i+1]-nums[i] == 1 {
			countSequence += 1

			if countSequence == 1 {
				str += strconv.Itoa(nums[i])
			}

		} else {

			if str != "" {
				str += "->" + strconv.Itoa(nums[i])
			} else {
				str += strconv.Itoa(nums[i])
			}

			result = append(result, str)
			countSequence = 0
			str = ""
		}

	}

	return result
}

// func summaryRanges(nums []int) []string {
// 	var (
// 		str string
// 		res []string
// 	)
//     if len(nums) == 0 {
//         return res
//     } else if len(nums) == 1 {
// 		str = fmt.Sprintf("%v", nums[0])
// 		res = append(res, str)
// 	}
// 	first, last := nums[0], 0

// 	for i := 1; i < len(nums); i++ {

// 		if math.Abs(float64(nums[i]) - float64(nums[i-1])) > 1 {
// 			last = nums[i-1]
// 			if last == first {
// 				str = fmt.Sprintf("%v", last)
// 				res = append(res, str)
// 			} else {
// 				str = fmt.Sprintf("%v->%v", first, last)
// 				res = append(res, str)
// 			}
// 			first = nums[i]
// 		}
//         if nums[i] == nums[len(nums)-1] {
// 			last = nums[i]
// 			if last == first {
// 				str = fmt.Sprintf("%v", last)
// 				res = append(res, str)
// 			} else {
// 				str = fmt.Sprintf("%v->%v", first, last)
// 				res = append(res, str)
// 			}
// 		}
// 	}

// 	return res
// }

// func main() {
// 	n := []int{0,2,3,4,6,8,9}

// 	fmt.Println(summaryRanges(n))
// }
