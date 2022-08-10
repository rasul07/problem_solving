package main

func taumBday(b int32, w int32, bc int32, wc int32, z int32) int64 {
	// Write your code here

	var (
		result       int64
		black, white uint64
	)

	if bc >= wc+z {
		black = uint64(b * (wc + z))
	} else {
		black = uint64(b * bc)
	}

	if wc >= bc+z {
		white = uint64(w * (bc + z))
	} else {
		white = uint64(w * wc)
	}

	result = int64(black + white)

	return result
}

// func main() {
// 	var b, w, bc, wc, z int32 = 58987449, 22313527, 461810, 182410, 378447

// 	fmt.Println(taumBday(b, w, bc, wc, z))
// }
