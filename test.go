package main

import (
	"fmt"
)

func main() {

	var x [9][9]int

	for i := range 9 {
		for j := range 9 {
			x[i][j] = 0
		}
	}
	fmt.Println(x)
}
