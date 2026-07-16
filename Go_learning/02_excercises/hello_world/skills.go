package main

import "fmt"

func main() {
	skills := make([]string, 0, 5)
	skills = append(skills, "Go")
	skills = append(skills, "Git")
	skills = append(skills, "Linux")
	fmt.Println(skills, len(skills), cap(skills))

}
