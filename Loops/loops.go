package main

import (
	"fmt"
)

func main() {
	// for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// while loop
	plantHeight := 1
	for plantHeight < 5 {
		fmt.Println("Still growing! current height:", plantHeight)
		plantHeight++
	}
	fmt.Println("plant has grown to ", plantHeight, " inches")

	// fizzbuzz
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else {
			fmt.Println(i)
		}
	}

	printPrimes(100)
}

func printPrimes(max int) {
	for i := 2; i < max+1; i++ {
		if i == 2 {
			fmt.Println(i)
			continue
		}
		if i%2 == 0 {
			continue
		}
		isPrime := true
		for j := 3; j*j < i+1; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Println(i)
		}
	}
}
