package main

func matchingStrings(strings []string, queries []string) []int32 {
	// Write your code here
	var result []int32
	for _, v := range queries {
		var count int32
		for i := 0; i < len(strings); i++ {
			if v == strings[i] {
				count += 1
			}
		}
		result = append(result, count)
	}
	return result
}
