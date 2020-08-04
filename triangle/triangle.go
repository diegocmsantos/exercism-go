package triangle

import "math"

// Kind Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

func (k Kind) isValidTriangle(a, b, c float64) bool {
	switch {
	case a <= 0 || math.IsNaN(a) || math.IsInf(a, 0):
		return false
	case b <= 0 || math.IsNaN(b) || math.IsInf(b, 0):
		return false
	case c <= 0 || math.IsNaN(c) || math.IsInf(c, 0):
		return false
	case a+b < c:
		return false
	case a+c < b:
		return false
	case b+c < a:
		return false
	default:
		return true
	}
}

func (k Kind) getType(a, b, c float64) Kind {
	if !k.isValidTriangle(a, b, c) {
		return NaT
	}
	if a == b && b == c && c == a {
		return Equ
	} else if a == b || b == c || c == a {
		return Iso
	}
	return Sca
}

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	return k.getType(a, b, c)
}
