package main

import "fmt"

type car struct {
	Make string
	Model string
	Height int
	Width int
	FrontWheel Wheel
	BackWheel Wheel
}

type Wheel struct {
	Radius int
	Material string
}

func main()  {
	myCar := car{}

	fmt.Println(myCar.FrontWheel)
}