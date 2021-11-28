package main

import "fmt"

func main() {

	var a [5]int
	for i := 0; i < 5; i++ {
		a[i] = i * 9
	}
	fmt.Println("emp:", a)

	var threeD [3][3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				threeD[i][j][k] = i + j + k
			}
		}
	}
	fmt.Println("3d: ", threeD)
}
