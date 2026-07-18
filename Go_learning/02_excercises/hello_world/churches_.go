package main

import "fmt"

func main() {
	churches := []string{
		"Deeper Life",
		"Pentecostal",
		"Redeemed",
	}
	fmt.Println(churches)
	fmt.Println(churches[len(churches)-1])
	fmt.Println(churches[len(churches)-2])

}
