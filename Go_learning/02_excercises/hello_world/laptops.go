package main

import "fmt"

func main() {
	laptops := []string{
		"HP",
		"Dell",
		"Asus",
	}
	laptops = append(laptops, "Macbook")
	laptops = append(laptops, "Chromebook")
	fmt.Println(laptops)
}
