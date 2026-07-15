package main

import "fmt"

func main() {
	var countries []string
	countries = append(countries, "Canada")
	countries = append(countries, "Nigeria")
	countries = append(countries, "Austrialia")
	fmt.Println(countries)
	fmt.Println(len(countries))
}
