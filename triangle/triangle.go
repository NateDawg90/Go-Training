package triangle

import (
	"fmt"
	"math"
)

const testVersion = 3

func KindFromSides(a, b, c float64) Kind {
	var kind Kind
	if a <= 0 || b <= 0 || c <= 0 {
		kind = NaT
	} else if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return NaT
	} else if math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		return NaT
	} else if a > b+c || b > a+c || c > a+b {
		kind = NaT
	} else if a == b && b == c {
		kind = Equ
	} else if a == b || b == c || a == c {
		kind = Iso
	} else {
		kind = Sca
	}

	fmt.Println(kind)
	return kind

}


type Kind string

const (
	NaT Kind = "Not a triangle" // not a triangle
	Equ Kind = "equilateral"    // equilateral
	Iso Kind = "Isoceles"       // isosceles
	Sca Kind = "scalene"        // scalene
)
