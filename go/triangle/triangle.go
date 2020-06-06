// Package triangle contains functions triangles
package triangle

import "math"

type Kind string

const (
	NaT = "NotATriangle"
	Equ = "Equilateral"
	Iso = "Isoceles"
	Sca = "Scalene"
)

// KindFromSides returns the Kind of triangle given it's side lengths
func KindFromSides(a, b, c float64) Kind {
	// sides of triangle must be non-zero, positive, and finite in length
	if !allPositiveAndFinite(a, b, c) {
		return NaT
	}

	// sides must not violate the triangle inequality
	if a+b < c || a+c < b || b+c < a {
		return NaT
	}

	// all sides are equal
	if a == b && b == c {
		return Equ
	}

	// two sides are equal
	if a == b || b == c || a == c {
		return Iso
	}

	// three sides of different lengths
	return Sca
}

// allPositiveAndFinite returns true if all arguments are non-zero, positive, and finite
func allPositiveAndFinite(nums ...float64) bool {
	for _, num := range nums {
		if math.IsNaN(num) || math.IsInf(num, 0) || num <= 0 {
			return false
		}
	}

	return true
}
