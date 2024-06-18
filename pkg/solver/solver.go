package solver

import (
	"fmt"
	"math"
	"vorto-challenge/pkg/model"
)

const (
	MaxDriveTime  = 12 * 60 // 12 hours in minutes
	DepotX        = 0.0
	DepotY        = 0.0
	CostPerDriver = 500.0
)

func distance(p1, p2 model.Point) float64 {
	return math.Sqrt((p2.X-p1.X)*(p2.X-p1.X) + (p2.Y-p1.Y)*(p2.Y-p1.Y))
}

func calculateLoadDistance(load model.Load) float64 {
	return distance(model.Point{DepotX, DepotY}, load.Pickup) + distance(load.Pickup, load.Dropoff) + distance(load.Dropoff, model.Point{DepotX, DepotY})
}

func calculateTotalDistance(route []int, loads []model.Load) float64 {
	totalDistance := 0.0
	currentLocation := model.Point{DepotX, DepotY}
	for _, loadID := range route {
		load := loads[loadID-1]
		totalDistance += distance(currentLocation, load.Pickup)
		totalDistance += distance(load.Pickup, load.Dropoff)
		currentLocation = load.Dropoff
	}
	totalDistance += distance(currentLocation, model.Point{DepotX, DepotY})
	return totalDistance
}

func nearestNeighbor(loads []model.Load, maxDriveTime float64) ([][]int, []bool) {
	numLoads := len(loads)
	visited := make([]bool, numLoads)
	var routes [][]int

	for i := 0; i < numLoads; i++ {
		if visited[i] {
			continue
		}
		route := []int{loads[i].ID}
		visited[i] = true
		currentLoad := loads[i]
		currentDriveTime := calculateLoadDistance(currentLoad)

		for {
			nextLoadIndex := -1
			minDistance := math.MaxFloat64

			for j := 0; j < numLoads; j++ {
				if visited[j] {
					continue
				}
				distanceToPickup := distance(currentLoad.Dropoff, loads[j].Pickup)
				if currentDriveTime+distanceToPickup+calculateLoadDistance(loads[j]) <= maxDriveTime {
					if distanceToPickup < minDistance {
						minDistance = distanceToPickup
						nextLoadIndex = j
					}
				}
			}

			if nextLoadIndex == -1 {
				break
			}

			route = append(route, loads[nextLoadIndex].ID)
			currentDriveTime += minDistance + calculateLoadDistance(loads[nextLoadIndex])
			currentLoad = loads[nextLoadIndex]
			visited[nextLoadIndex] = true
		}

		routes = append(routes, route)
	}

	return routes, visited
}

func twoOpt(route []int, loads []model.Load) []int {
	best := route
	improved := true

	for improved {
		improved = false
		for i := 1; i < len(route)-1; i++ {
			for j := i + 1; j < len(route); j++ {
				if j-i == 1 {
					continue
				}
				newRoute := make([]int, len(route))
				copy(newRoute, route)
				reverse(newRoute, i, j)
				if calculateTotalDistance(newRoute, loads) < calculateTotalDistance(best, loads) {
					best = newRoute
					improved = true
				}
			}
		}
		route = best
	}
	return best
}

func reverse(route []int, i, j int) {
	for k := 0; k < (j-i+1)/2; k++ {
		route[i+k], route[j-k] = route[j-k], route[i+k]
	}
}

func mergeRoutes(routes [][]int, loads []model.Load, maxDriveTime float64) [][]int {
	merged := [][]int{}

	for len(routes) > 0 {
		route := routes[0]
		routes = routes[1:]

		for i := 0; i < len(routes); i++ {
			newRoute := append(route, routes[i]...)
			if calculateTotalDistance(newRoute, loads) <= maxDriveTime {
				route = newRoute
				routes = append(routes[:i], routes[i+1:]...)
				i--
			}
		}

		merged = append(merged, route)
	}

	return merged
}

func SolveVRP(loads []model.Load) [][]int {
	// print how many loads are there
	fmt.Println("Number of loads: ", len(loads))

	routes, visited := nearestNeighbor(loads, MaxDriveTime)
	for i := range routes {
		routes[i] = twoOpt(routes[i], loads)
	}
	routes = mergeRoutes(routes, loads, MaxDriveTime)

	// Ensure all loads are assigned
	unassignedLoads := []int{}
	for i, load := range loads {
		if !visited[i] {
			unassignedLoads = append(unassignedLoads, load.ID)
		}
	}

	// print how many unassigned loads are there
	fmt.Println("Number of unassigned loads: ", len(unassignedLoads))

	for _, loadID := range unassignedLoads {
		routes = append(routes, []int{loadID})
	}

	// print how many routes are there
	fmt.Println("Number of routes: ", len(routes))

	// Debugging: Check the total distance and load for each route
	for i, route := range routes {
		totalDistance := calculateTotalDistance(route, loads)
		fmt.Printf("Route %d: %v, Total Distance: %f\n", i+1, route, totalDistance)
	}

	return routes
}
