package main

import (
	"bufio"
	"fmt"
	"os"
	"vorto-challenge/pkg/model"
	"vorto-challenge/pkg/solver"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: vrp <path_to_problem>")
		return
	}

	problemPath := os.Args[1]

	file, err := os.Open(problemPath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	var loads []model.Load
	scanner := bufio.NewScanner(file)
	// Skip the header
	scanner.Scan()

	for scanner.Scan() {
		var load model.Load
		_, err := fmt.Sscanf(scanner.Text(), "%d (%f,%f) (%f,%f)", &load.ID, &load.Pickup.X, &load.Pickup.Y, &load.Dropoff.X, &load.Dropoff.Y)
		if err != nil {
			fmt.Printf("Error parsing line: %v\n", err)
			return
		}
		loads = append(loads, load)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	solution := solver.SolveVRP(loads)
	for _, driver := range solution {
		fmt.Printf("%v\n", driver)
	}
}
