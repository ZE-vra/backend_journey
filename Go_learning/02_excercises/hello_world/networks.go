package main

import "fmt"

func main() {
	networks := make(map[string]string, 4)
	networks["MTN"] = "1st"
	networks["Airtel"] = "2nd"
	networks["Glo"] = "3rd"
	networks["9mobile"] = "4th"
	best_network, ok := networks["Glo"]
	fmt.Println(best_network, ok)
}
