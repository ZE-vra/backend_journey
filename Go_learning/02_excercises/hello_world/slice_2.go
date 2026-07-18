package main

import "fmt"

func main() {
	x := []int{5, 10, 15}
	y := make([]int, 3)
	y = x
	y[1] = 13
	fmt.Println(x)
	fmt.Println(y)

}
