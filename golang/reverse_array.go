package main

func reverseArray(a []int32) []int32 {
    // Write your code here
    var reverseArray []int32
    for i := len(a) - 1; i >= 0; i-- {
        reverseArray = append(reverseArray, a[i])
    }
    return reverseArray
}
