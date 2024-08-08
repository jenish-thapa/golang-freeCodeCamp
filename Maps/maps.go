package main

import "fmt"

func main() {
	// define a map
	ages := make(map[string]int)
	ages["John"] = 37
	ages["Mary"] = 21

	// define and initialize a map
	ages2 := map[string]int{
		"Jenish": 100,
		"Jay": 50,
	}

	// delete a key
	delete(ages, "Mary") 

	// check if a key exists
	ele, ok := ages2["Jenish"]

	fmt.Println(ele, ok)

	fmt.Println("Ages:", ages)
	fmt.Println("Ages:", ages2)
}
