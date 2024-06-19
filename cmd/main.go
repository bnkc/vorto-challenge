package main

import (
	"fmt"
	"os"
	"vorto-challenge/internal"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./vorto <problem_path>")
		return
	}

	problemPath := os.Args[1]

	// Plan the routes for the VRP problem
	if err := internal.PlanRoutes(problemPath); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
