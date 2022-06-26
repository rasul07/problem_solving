package main

import "fmt"

// func jumpingOnClouds(c []int32) int32 {
//     // Write your code here
//     var (
//         count int32
//         l int = len(c)-1
//     )

//     for i := 0; i < l; i += 2 {
//         if i+2 < len(c){
//             if c[i+1] == 1 {
//                 count++
//             } else if c[i+2] == 1 {
//                 count += 2
//             } else {
//                 count++
//             }
//         }  else {
//             count++
//         }
//     }

//     return count
// }

func jumpingOnClouds(c []int32) int32 {
	// Write your code here
	var (
		count int32
	)

	for i := 0; i < len(c) - 1; i ++ {
		if c[i]==0{
			i++
		}
			count++
		}

	return count
}

func main() {
	c := []int32{0, 1, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 1, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0}

	fmt.Println(jumpingOnClouds(c))
}

