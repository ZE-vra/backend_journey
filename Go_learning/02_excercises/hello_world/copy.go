package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}
	y := make([]int, 4)
	copy(y, x)
	y = append(y, 25)
	fmt.Println(x, y)
}
