package darts

import "math"

// Score calculates points earned by a dart landing at a given point
func Score(x, y float64) int {
	distance := math.Sqrt(x*x + y*y)
	switch {
	case distance <= 1:
		return 10
	case distance <= 5:
		return 5
	case distance <= 10:
		return 1
	default:
		return 0
	}
}
