package classes

import (
	"fmt"

	"github.com/gutmanndev/go-basic-platzi-course/classes/challenges"
)

func ThirdClass() {
	// Algorithmyc operations Class
	// declaring some variables to do math with them
	x := 10
	y := 20

	// Sum
	res := x + y

	// substract
	res = x - y

	// multiplication
	res = x * y

	// Division
	res = x / y

	// Module
	res = x % y

	// Incremental: sum 1 to a number (used in for loops)
	x++
	fmt.Println(x)
	// Decremental
	x--
	fmt.Println(x)

	fmt.Println(res)

	fmt.Println(challenges.CalculateArea("circle"))
}
