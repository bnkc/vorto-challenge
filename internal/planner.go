package internal

import (
	"fmt"
)

const maxDriveTime = 12 * 60 // 12 hours in minutes

// PlanRoutes reads the problem from a file and plans the routes for the VRP problem.
func PlanRoutes(filePath string) error {
	loads, err := ParseInput(filePath)
	if err != nil {
		return fmt.Errorf("failed to parse input: %w", err)
	}

	visited := make(map[int]bool)
	var drivers []*Driver
	depot := Point{X: 0, Y: 0}
	driverID := 1

	for len(visited) < len(loads) {
		currentPoint := depot
		currentDriveTime := 0.0
		newDriver := NewDriver(driverID)

		for {
			ld, found := nearestLoad(currentPoint, loads, visited)
			if !found || !newDriver.AddLoad(ld, currentPoint, maxDriveTime) {
				break
			}
			visited[ld.ID] = true
			currentPoint = ld.Dropoff
		}
		// Return to depot
		currentDriveTime += Distance(currentPoint, depot)
		if currentDriveTime <= maxDriveTime {
			drivers = append(drivers, newDriver)
			driverID++
		}
	}

	for _, d := range drivers {
		fmt.Println(d)
	}

	return nil
}

// nearestLoad finds the nearest unvisited load to the current point.
func nearestLoad(current Point, loads []Load, visited map[int]bool) (Load, bool) {
	var nearest Load
	minDistance := float64(maxDriveTime)
	found := false

	for _, ld := range loads {
		if visited[ld.ID] {
			continue
		}
		distToPickup := Distance(current, ld.Pickup)
		if distToPickup < minDistance {
			minDistance = distToPickup
			nearest = ld
			found = true
		}
	}

	return nearest, found
}
