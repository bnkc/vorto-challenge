package internal

import (
	"fmt"
	"strings"
)

type Driver struct {
	ID        int
	Loads     []Load
	DriveTime float64
}

func NewDriver(id int) *Driver {
	return &Driver{ID: id}
}

// AddLoad adds a load to the driver's schedule if it doesn't exceed the max drive time.
func (d *Driver) AddLoad(ld Load, currentPoint Point, maxDriveTime float64) bool {
	additionalDriveTime := EuclideanDistance(currentPoint, ld.Pickup) + ld.Distance + EuclideanDistance(ld.Dropoff, Point{X: 0, Y: 0})
	if d.DriveTime+additionalDriveTime <= maxDriveTime {
		d.Loads = append(d.Loads, ld)
		d.DriveTime += EuclideanDistance(currentPoint, ld.Pickup) + ld.Distance
		return true
	}
	return false
}

func (d *Driver) String() string {
	var loadIDs []string
	for _, ld := range d.Loads {
		loadIDs = append(loadIDs, fmt.Sprintf("%d", ld.ID))
	}
	return fmt.Sprintf("[%s]", strings.Join(loadIDs, ","))
}
