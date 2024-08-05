package main

import "fmt"

// func declaration
func concat(s1 string, s2 string) string {
	return s1 + s2
}

// short-cut to define args with same type
func add(x, y int) int {
	return x + y
}

// explicit return 
func getPoint() (int, int) {
	return 3, 4
}

// implicit return
func getCoords() (x, y int) {
	return
}

func main() {
	fmt.Println(concat("Hello", " World!"))

	// ignoring return vals
	x, _ := getPoint()

	x1, _ := getCoords()

	fmt.Println(x, x1)
}