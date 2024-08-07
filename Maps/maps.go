package main

import "fmt"

func main() {
	ages := make(map[string]int)
	ages["John"] = 37
	ages["Mary"] = 21

	ages2 := map[string]int{
		"Jenish": 100,
		"Jay": 50,
	}

	fmt.Println("Ages:", ages)
	fmt.Println("Ages:", ages2)
}
