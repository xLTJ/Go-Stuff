package triangle

type Kind int

const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

// KindFromSides returns the type of triangle based on the provided site lengths
func KindFromSides(a, b, c float64) Kind {
	if a <= 0 || b <= 0 || c <= 0 {
		return NaT
	}

	switch {
	case a+b <= c || b+c <= a || a+c <= b:
		return NaT
	case a == b && a == c:
		return Equ
	case a == b || a == c || b == c:
		return Iso
	default:
		return Sca

	}
}
