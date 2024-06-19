package internal

import (
	"fmt"
)

const maxDriveTime = 12 * 60 // 12 hours in minutes

// PlanRoutes reads the problem from a file and plans route for the challenge
func PlanRoutes(filePath string) error {
	loads, err := ParseInput(filePath)
	if err != nil {
		return fmt.Errorf("failed to parse: %w", err)
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
			ld, found := nearestNeighbor(currentPoint, loads, visited)
			// If we can't find any more loads or the driver can't add the load, break
			if !found || !newDriver.AddLoad(ld, currentPoint, maxDriveTime) {
				break
			}
			visited[ld.ID] = true
			currentPoint = ld.Dropoff
		}
		// Return to depot
		currentDriveTime += EuclideanDistance(currentPoint, depot)
		if currentDriveTime <= maxDriveTime {
			drivers = append(drivers, newDriver)
			driverID++
		}
	}

	// Apply Two-Opt algo to driver's route
	for _, d := range drivers {
		twoOpt(d, depot)
	}

	for _, d := range drivers {
		fmt.Println(d)
	}

	return nil
}

// nearestNeighbor finds the nearest unvisited load to the current point.
func nearestNeighbor(current Point, loads []Load, visited map[int]bool) (Load, bool) {
	var nearest Load
	minDistance := float64(maxDriveTime)
	found := false

	for _, ld := range loads {
		if visited[ld.ID] {
			continue
		}
		distToPickup := EuclideanDistance(current, ld.Pickup)
		if distToPickup < minDistance {
			minDistance = distToPickup
			nearest = ld
			found = true
		}
	}

	return nearest, found
}

// twoOpt applies the Two-Opt heuristic to try to improve the driver's route. if it's possible.
func twoOpt(driver *Driver, depot Point) {
	for {
		improved := false
		for i := 1; i < len(driver.Loads)-1; i++ {
			for j := i + 1; j < len(driver.Loads); j++ {
				if swap(driver, i, j, depot) {
					improved = true
				}
			}
		}
		if !improved {
			break
		}
	}
}

func swap(driver *Driver, i, j int, depot Point) bool {
	oldDist := totalDistance(driver, depot)
	reverse(driver.Loads, i, j)
	newDist := totalDistance(driver, depot)
	if newDist < oldDist {
		driver.DriveTime -= oldDist - newDist
		return true
	}
	// Swap back if not improved
	reverse(driver.Loads, i, j)
	return false
}

func reverse(loads []Load, i, j int) {
	for k := 0; k < (j-i+1)/2; k++ {
		loads[i+k], loads[j-k] = loads[j-k], loads[i+k]
	}
}

// calculates the total distance of the driver's route + the return to depot
func totalDistance(driver *Driver, depot Point) float64 {
	totalDist := 0.0
	currentPoint := depot
	for _, ld := range driver.Loads {
		totalDist += EuclideanDistance(currentPoint, ld.Pickup) + ld.Distance
		currentPoint = ld.Dropoff
	}
	totalDist += EuclideanDistance(currentPoint, depot)
	return totalDist
}
