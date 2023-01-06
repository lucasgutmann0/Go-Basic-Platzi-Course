package second_class

import "fmt"

func SecondClass() {
	// HOW TO DECLARE CONSTANTS!!
	const pi_float float64 = 3.14 // this is one way, specifying the type of it
	// this it other way, without specifying the type
	const pi_untyped = 3.1416 // this const shows it self as untyped float

	fmt.Println("pi_typed:", pi_float)
	fmt.Println("pi_untyped:", pi_untyped)

	// HOW TO DECLARE VARIABLES!!

	// using a colon before the equal say to go
	// that the variable hasn't been declared before and you're declaring it
	// and also initializing it
	base := 32
	var altura int = 14 // declaring and initializing it and specyfing the type
	var area int        // just declaring it

	// using it to avoid erros on go
	fmt.Println(base, altura, area)

	// WHAT ARE ZERO VALUES
	// most programming languages use null or none
	// as the default value for a none initialized variable
	// but in go there are zero values

	// this are the zero values for each type:
	var a int     // 0
	var b float64 // 0
	var c string  // ""
	var d bool    // false

	// using it to avoid erros on go
	fmt.Println(a, b, c, d)

	// LET'S CALCULATE THE AREA OF AN SQUARE
	const squareBase int = 10
	squareArea := squareBase * squareBase

	fmt.Println("square_area: ", squareArea)
}
