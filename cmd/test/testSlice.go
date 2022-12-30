package main

import (
	"fmt"
)

func main() {
	x := []int64{1, 2, 3, 4, 5, 6}
	for i := 0; i < len(x)-1; i++ {
		if x[i] < x[i+1] {
			a := x[i]
			x[i] = x[i+1]
			x[i+1] = a
		}
	}
	fmt.Println(x)
}
