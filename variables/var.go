package main

import (
	"fmt"
	"math/cmplx"
)

func main()  {

	// var a bool
	// a = true

	// var a int
	// a = 15

	// var a float32
	// a = 15.0

	// var a string
	// a = "string example"

	// declare & initialize
	// a := 15

	// complex data type
	var a complex128
	a = cmplx.Acos(-2 + 12i)

	fmt.Println("value of a:", a ) 
}