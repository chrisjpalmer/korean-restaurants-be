package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/chrisjpalmer/korean-restaurants-be/internal/restaurant"
)

type Server struct {
	srv    *http.Server
	restdb restaurant.Database
}

func New(restdb restaurant.Database, port string) *Server {
	s := Server{
		restdb: restdb,
	}

	mux := http.NewServeMux()
	mux.Handle("/restaurant", http.HandlerFunc(s.FindRestaurants))

	s.srv = &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	return &s
}

func (s *Server) FindRestaurants(r http.ResponseWriter, rq *http.Request) {
	r.Header().Set("Access-Control-Allow-Origin", "*")

	// parse request
	nearby, err := parseNearby(rq.URL.Query().Get("nearby"))
	if err != nil {
		errorStatus(r, http.StatusBadRequest, err)
		return
	}
	withinMeters, err := parseWithinMeters(rq.URL.Query().Get("within_meters"))
	if err != nil {
		errorStatus(r, http.StatusBadRequest, err)
		return
	}

	// find restaurants
	rr, err := s.restdb.FindRestaurants(rq.Context(), nearby, withinMeters)
	if err != nil {
		log.Println(err)
		errorStatus(r, http.StatusInternalServerError, fmt.Errorf("an internal error occurred"))
		return
	}

	// response
	log.Printf("server: found %d restaurants", len(rr))
	rs := FindRestaurantsResponse{
		Restaurants: mapRestaurants(rr),
	}
	log.Printf("server: mapped %d restaurants", len(rs.Restaurants))
	successStatus(r, rs)
}

func errorStatus(r http.ResponseWriter, statusCode int, err error) {
	r.WriteHeader(statusCode)
	var rs ErrorResponse
	rs.Error = err.Error()
	b, err := json.Marshal(rs)
	if err != nil {
		log.Printf("error when marshalling error status: %v", err)
		return
	}
	_, err = r.Write(b)
	if err != nil {
		log.Printf("error when sending error status: %v", err)
		return
	}
}

func successStatus(r http.ResponseWriter, rs interface{}) {
	r.WriteHeader(200)
	b, err := json.Marshal(rs)
	if err != nil {
		log.Printf("error when marshalling response object: %v", err)
		return
	}
	_, err = r.Write(b)
	if err != nil {
		log.Printf("error when sending response: %v", err)
		return
	}
}

func parseNearby(s string) (restaurant.Coordinates, error) {
	ss := strings.Split(s, ",")
	if len(ss) != 2 {
		return restaurant.Coordinates{}, fmt.Errorf("invalid coordinate format: expected lon,lat")
	}
	lons := ss[0]
	lats := ss[1]
	lat, err := strconv.ParseFloat(lats, 64)
	if err != nil {
		return restaurant.Coordinates{}, fmt.Errorf("invalid coordinate format: invalid latitude: %w", err)
	}
	lon, err := strconv.ParseFloat(lons, 64)
	if err != nil {
		return restaurant.Coordinates{}, fmt.Errorf("invalid coordinate format: invalid longitude: %w", err)
	}
	return restaurant.Coordinates{
		Lat: lat,
		Lon: lon,
	}, nil
}

func parseWithinMeters(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

func mapRestaurants(rr []restaurant.Restaurant) []Restaurant {
	out := make([]Restaurant, len(rr))
	for i := range rr {
		out[i] = Restaurant{
			Name:        rr[i].Name,
			Description: rr[i].Description,
			Coordinates: Coordinates(rr[i].Coordinates),
		}
	}
	return out
}

func (s *Server) Serve() error {
	err := s.srv.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func (s *Server) Close() error {
	err := s.srv.Close()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}
