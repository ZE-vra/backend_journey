package main

import (
	"fmt"
)

func main() {
	var myDetails []string
	myDetails = append(myDetails, "Miracle")
	myDetails = append(myDetails, "Zevra")
	myDetails = append(myDetails, "Backend Engineer")
	fmt.Println(myDetails)
	fmt.Println(len(myDetails))
}
