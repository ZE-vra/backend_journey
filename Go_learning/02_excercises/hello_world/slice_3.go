package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}
	y := x
	fmt.Println(x)
	fmt.Println(y)
	y[2] = 7
	y[3] = 9
	fmt.Println(x)
	fmt.Println(y)
}
