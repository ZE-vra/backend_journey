package main

import "fmt"

func main() {
	x := 20
	y := 3.5
	z := float64(x) * y
	fmt.Println(z)
	fmt.Println(x + int(y))
	fmt.Println(float64(x) - y)
}
