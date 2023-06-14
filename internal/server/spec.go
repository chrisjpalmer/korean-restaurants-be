package server

type ErrorResponse struct {
	Error string `json:"error"`
}

type FindRestaurantsResponse struct {
	Restaurants []Restaurant `json:"restaurants"`
}

type Restaurant struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Coordinates Coordinates `json:"coordinates"`
}

type Coordinates struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}
