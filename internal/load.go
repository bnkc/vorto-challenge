package internal

// Load reps a load with an ID, pickup and dropoff points + the distance between them
type Load struct {
	ID       int
	Pickup   Point
	Dropoff  Point
	Distance float64
}

func NewLoad(id int, pickup, dropoff Point) Load {
	return Load{
		ID:       id,
		Pickup:   pickup,
		Dropoff:  dropoff,
		Distance: EuclideanDistance(pickup, dropoff),
	}
}
