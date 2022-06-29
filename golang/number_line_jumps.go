package main

import "fmt"

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
    if v2<v1 && (x1 - x2) % (v2 - v1) == 0 {
        return "YES"
    }
    return "NO"
}


func main() {
	x1, v1, x2, v2 := 6644, 5868, 8349, 3477
	
	fmt.Println(kangaroo(int32(x1), int32(v1), int32(x2), int32(v2)))
}
