package main

import "fmt"

func main() {
	scores := map[string]int{}
	scores["Zevra"] = 98
	scores["Sarah"] = 88
	scores["John"] = 76
	fmt.Println(scores, len(scores))
}
