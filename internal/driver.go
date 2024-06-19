package internal

import (
	"fmt"
	"strings"
)

// Depot location
// const Depot = Point{X: 0, Y: 0}

// Driver reps a driver with an ID, assigned loads, and total drive time
type Driver struct {
	ID        int
	Loads     []Load
	DriveTime float64
}

// NewDriver creates a new driver with the given ID
func NewDriver(id int) *Driver {
	return &Driver{ID: id}
}

// AddLoad adds a load to the driver's schedule if it doesn't exceed the max drive time.
func (d *Driver) AddLoad(ld Load, currentPoint Point, maxDriveTime float64) bool {
	additionalDriveTime := Distance(currentPoint, ld.Pickup) + ld.Distance + Distance(ld.Dropoff, Point{X: 0, Y: 0})
	if d.DriveTime+additionalDriveTime <= maxDriveTime {
		d.Loads = append(d.Loads, ld)
		d.DriveTime += Distance(currentPoint, ld.Pickup) + ld.Distance
		return true
	}
	return false
}

// String returns the driver's load schedule in the required format.
func (d *Driver) String() string {
	var loadIDs []string
	for _, ld := range d.Loads {
		loadIDs = append(loadIDs, fmt.Sprintf("%d", ld.ID))
	}
	return fmt.Sprintf("[%s]", strings.Join(loadIDs, ","))
}
