package main

import "fmt"

func main() {
	name := "Chosen"
	age := 24
	name = "Zevra"

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(name == "Zevra", "Hello"+" "+name)
	fmt.Println(name + " " + "is" + " " + "beautiful")
}
