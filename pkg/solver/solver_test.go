package solver

import (
	"testing"
	"vorto-challenge/pkg/model"
)

func TestSolveVRP(t *testing.T) {
	loads := []model.Load{
		{ID: 1, Pickup: model.Point{-9.1, -48.89}, Dropoff: model.Point{-116.78, 76.8}},
		{ID: 2, Pickup: model.Point{73.39, -86.93}, Dropoff: model.Point{-57.59, 28.66}},
		{ID: 3, Pickup: model.Point{-109.23, -94.63}, Dropoff: model.Point{134.99, -41.03}},
		{ID: 4, Pickup: model.Point{-126.03, 12.04}, Dropoff: model.Point{-102.91, -41.30}},
		{ID: 5, Pickup: model.Point{-113.02, -28.79}, Dropoff: model.Point{-5.19, -89.13}},
	}

	expectedDrivers := 3

	result := SolveVRP(loads)

	if len(result) != expectedDrivers {
		t.Fatalf("Expected %d drivers, got %d", expectedDrivers, len(result))
	}

	// Additional checks for the correctness of the routes
	for _, route := range result {
		totalDistance := calculateTotalDistance(route, loads)
		if totalDistance > MaxDriveTime {
			t.Fatalf("Route %v exceeds the maximum drive time", route)
		}
	}
}
