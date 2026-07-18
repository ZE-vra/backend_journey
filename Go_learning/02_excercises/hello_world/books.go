package main

import "fmt"

func main() {
	books := map[string]int{
		"Go":     5,
		"Python": 3,
		"Linux":  8,
	}
	copies, ok := books["Go"]
	fmt.Println(copies, ok)
	copies, ok = books["Rust"]
	fmt.Println(copies, ok)

}
