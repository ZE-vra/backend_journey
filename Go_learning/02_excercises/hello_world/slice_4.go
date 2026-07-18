package main

import "fmt"

func main() {
	x := []int{2, 4, 6, 8}
	y := x[:3]
	y[1] = 3
	fmt.Println(x)
	fmt.Println(y)

}
