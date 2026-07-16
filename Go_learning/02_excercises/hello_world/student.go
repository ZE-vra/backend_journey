package main

import "fmt"

func main() {
	profile := []string{
		"Chekwube",
		"Zevra",
	}
	fmt.Print(profile, len(profile), cap(profile))
	profile = append(profile, "Backend Engineer")
	fmt.Println(profile, len(profile), cap(profile))
}
