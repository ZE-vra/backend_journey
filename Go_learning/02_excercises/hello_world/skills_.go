package main

import "fmt"

func main() {
	skills := make([]string, 0, 4)
	fmt.Println(skills, len(skills), cap(skills))
	skills = append(skills, "Go", "Backend", "Git", "Linux", "Docker")
	fmt.Println(skills, len(skills), cap(skills))
}
