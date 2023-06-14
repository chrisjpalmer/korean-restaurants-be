package restaurant

import "context"

type Database interface {
	FindRestaurants(ctx context.Context, center Coordinates, radiusMeters float64) ([]Restaurant, error)
}

type Restaurant struct {
	Name        string
	Description string
	Coordinates Coordinates
}

type Coordinates struct {
	Lat float64
	Lon float64
}
