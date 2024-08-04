package main

import "fmt"

func main() {
	// declaring vars
	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool
	var username string
	fmt.Printf(
		"%v %f %v %q\n",
		smsSendingLimit,
		costPerSMS,
		hasPermission,
		username,
	)

	// short var declaration
	wish := "happy birthday"
	fmt.Println(wish)

	// get type of a var
	fmt.Printf("The type of costPerSMS is %T\n", costPerSMS)

	// multiple var declaration
	a , b := 1, "hello"
	fmt.Println(a, b)

	// type casting
	balance := 10.9
	fmt.Printf("The type of balance is %T\n", int(balance))

	// constants
	const PI = 3.14

	// computed constants
	const fName = "Jenish"
	const lName = "Thapa"
	const fullName = fName + " " + lName

	// formating strings
	const name = "Jenish"
	const openRate = 30.59079

	msg := fmt.Sprintf("Hi %s, your open rate is %.2f percent", name, openRate,)
	fmt.Println(msg)
}