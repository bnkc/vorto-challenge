package internal

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// ParseInput reads the loads from the input file.
func ParseInput(filePath string) ([]Load, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var loads []Load
	scanner := bufio.NewScanner(file)
	scanner.Scan() // skip header
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		id, _ := strconv.Atoi(fields[0])
		pickup := parsePoint(fields[1])
		dropoff := parsePoint(fields[2])
		loads = append(loads, NewLoad(id, pickup, dropoff))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return loads, nil
}

// parsePoint parses a string representation of a point into a Point struct.
func parsePoint(s string) Point {
	s = strings.Trim(s, "()")
	coords := strings.Split(s, ",")
	x := parseFloat(coords[0])
	y := parseFloat(coords[1])
	return Point{X: x, Y: y}
}

// parseFloat converts a string to a float64.
func parseFloat(s string) float64 {
	val, _ := strconv.ParseFloat(s, 64)
	return val
}
