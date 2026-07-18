package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}
	y := x[1:3]
	y[1] = 6
	fmt.Println(x, y)
	x = append(x, 25)
	fmt.Println(y[1], x)
}
