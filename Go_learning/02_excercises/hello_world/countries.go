package main

import "fmt"

func main() {
	var countries []string
	countries = append(countries, "Canada")
	countries = append(countries, "Nigeria")
	countries = append(countries, "Austrialia")
	fmt.Println(countries)
	fmt.Println(len(countries))
	fmt.Println(countries[2])
	countries[1] = "Argentina"
	countries[1] = "Spain"
	fmt.Print(countries)
	fmt.Print(countries[len(countries)])
}
