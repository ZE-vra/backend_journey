package main

import "fmt"

func main() {
	friends := []string{
		"Daniella",
		"Timi",
	}
	fmt.Println(friends)
	friends = append(friends, "Zaram")
	friends = append(friends, "Francisca")
	fmt.Println(friends)
	fmt.Println(len(friends))
}
