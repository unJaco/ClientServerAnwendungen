package model

// eine location hat latitude und longitude
type Location struct {
	Lat 			float64 		`json:"lat"`
	Lon 			float64			`json:"lon"`
}