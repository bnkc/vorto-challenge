package internal

import "math"

// Represents a Cartesian point.
type Point struct {
	X float64
	Y float64
}

func EuclideanDistance(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(p2.X-p1.X, 2) + math.Pow(p2.Y-p1.Y, 2))
}
