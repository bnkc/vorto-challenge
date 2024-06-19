package internal

import "math"

// Point represents a Cartesian point.
type Point struct {
	X float64
	Y float64
}

// Distance calculates the Euclidean distance between two points.
func Distance(p1 Point, p2 Point) float64 {
	return math.Sqrt(math.Pow(p2.X-p1.X, 2) + math.Pow(p2.Y-p1.Y, 2))
}
