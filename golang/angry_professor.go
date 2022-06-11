package main

func angryProfessor(k int32, a []int32) string {
	// Write your code here
	var presentStd int32
	for _, v := range a {
		if v <= 0 {
			presentStd += 1
		}
	}
	if presentStd >= k {
		return "NO"
	} else {
		return "YES"
	}

}
