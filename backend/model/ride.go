package model

import (
	enums "github.com/unJaco/UberClientServer/backend/model/enums"
	"gorm.io/gorm"
)

// eine Fahrt hat folgende Struktur
type Ride struct {
	// gorm.Model fügt automatisch eine ID, createdAt, updatedAt und deletedAt zur Tablle hinzu, bei Automigration
	gorm.Model
	VehicleId  uint             `json:"vehicleId"`
	DriverId   uint             `json:"driverId"`
	CustomerId uint             `json:"customerId"`
	StartLon   float64          `json:"startLon"`
	StartLat   float64          `json:"startLat"`
	EndLon     float64          `json:"endLon"`
	EndLat     float64          `json:"endLat"`
	Distance   float64          `json:"distance"`
	Status     enums.RideStatus `json:"status"`
	Price      float32          `json:"price"`
	Rating     uint				`json:"rating"`
}
