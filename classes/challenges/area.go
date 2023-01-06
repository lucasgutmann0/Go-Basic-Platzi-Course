package challenges

import (
	"fmt"
	"math"
	"strings"
)

func CalculateArea(opt string) (area_value float64, opt_type string) {
	// lower option
	lowered_opt := strings.ToLower(opt)
	opt_type = lowered_opt
	// message
	fmt.Println("Please type the requested values: ")
	// Check cases
	switch lowered_opt {
	case "rectangle":
		var b, h float64
		// base
		fmt.Print("Base: ")
		// base input
		fmt.Scan(&b)

		// height
		fmt.Print("Height: ")
		// height input
		fmt.Scan(&h)
		area_value = b * h
	case "circle":
		var r float64
		// base
		fmt.Print("Radius: ")
		// base input
		fmt.Scan(&r)

		area_value = 3.1416 * math.Pow(r, 2) // r^2
	case "trapeze":
		var b1, b2, h float64
		// base 1
		fmt.Print("Base 1: ")
		// base 1 input
		fmt.Scan(&b1)

		// base 2
		fmt.Print("Base 2: ")
		// base 2 input
		fmt.Scan(&b2)

		// height
		fmt.Print("Height: ")
		// height input
		fmt.Scan(&h)
		area_value = ((b1 + b2) / 2) * h
	default:
		area_value = -1
		opt_type = "nothing"
	}

	return area_value, opt_type
}
