package model

type Point struct {
	X, Y float64
}

type Load struct {
	ID      int
	Pickup  Point
	Dropoff Point
}

type Driver struct {
	Loads []int
}
