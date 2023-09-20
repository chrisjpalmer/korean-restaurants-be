package restdb

import (
	"context"
	"fmt"
	"regexp"
	"strconv"

	"github.com/chrisjpalmer/korean-restaurants-be/internal/restaurant"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	pool *pgxpool.Pool
}

func New(connString string) (*Database, error) {
	cfg, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}
	p, err := pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		return nil, err
	}
	return &Database{
		pool: p,
	}, nil
}

func (d *Database) FindRestaurants(ctx context.Context, center restaurant.Coordinates, radiusMeters float64) ([]restaurant.Restaurant, error) {
	// query database
	// NOTE: qry is sql injection safe
	qry := findRestaurantsQuery(center)
	rows, err := d.pool.Query(ctx, qry)
	if err != nil {
		return nil, err
	}

	// map results
	rr := make([]restaurant.Restaurant, 0)
	for rows.Next() {
		var r restaurant.Restaurant
		var dist float64
		var pointStr string
		rows.Scan(&r.Name, &r.Description, &dist, &pointStr)
		r.Coordinates, err = mapPointString(pointStr)
		if err != nil {
			return nil, err
		}
		rr = append(rr, r)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return rr, nil
}

// findRestauarantsQuery - NOTE: this is sql injection safe
func findRestaurantsQuery(center restaurant.Coordinates) string {
	//NOTE: pointStr is sql injection safe.
	pointStr := mapCoordinate(center)
	qry := `
		SELECT 
			kr.name,
			kr.description, 
			kr.the_geog <-> %s AS dist,
			ST_AsText(kr.the_geog)
		FROM
			korean_restaurants kr
		WHERE kr.the_geog <-> %s < $1
		ORDER BY dist
		LIMIT 3;
	`

	// this is safe because all inputs are safe
	qry = fmt.Sprintf(qry, pointStr, pointStr)

	return qry
}

// mapCoordinate - NOTE: this method is sql injection safe
func mapCoordinate(c restaurant.Coordinates) string {
	return fmt.Sprintf(`'SRID=4326;POINT(%.14f %.14f)'::geography`, c.Lon, c.Lat)
}

// POINT Z (126.9775201550173 37.22450239990378 35.55338264945428)
var pointRegex = regexp.MustCompile(`POINT Z \(([0-9\.]*) ([0-9\.]*).*\)`)

func mapPointString(s string) (restaurant.Coordinates, error) {
	strs := pointRegex.FindAllStringSubmatch(s, -1)
	if len(strs) != 1 {
		return restaurant.Coordinates{}, fmt.Errorf("mapPointString: expected 1 match but got %d", len(strs))
	}
	if len(strs[0]) != 3 {
		return restaurant.Coordinates{}, fmt.Errorf("mapPointString: expected 3 submatches match but got %d", len(strs))
	}
	lons := strs[0][1]
	lats := strs[0][2]
	lon, err := strconv.ParseFloat(lons, 64)
	if err != nil {
		return restaurant.Coordinates{}, fmt.Errorf("mapPointString: could not convert lon: %w", err)
	}
	lat, err := strconv.ParseFloat(lats, 64)
	if err != nil {
		return restaurant.Coordinates{}, fmt.Errorf("mapPointString: could not convert lat: %w", err)
	}
	return restaurant.Coordinates{
		Lat: lat,
		Lon: lon,
	}, nil
}
