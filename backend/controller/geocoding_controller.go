package controller

import (
	"math"
)

const earthRadius float64 = 6371 // Erdradius in KM

type GeoCodingController struct {}


// toRadians ist eine Funktion, die Grad in Radiant umrechnet
func toRadians(degrees float64) float64 {
	return degrees * math.Pi / 180  
}

// toDegrees ist eine Funktion, die Radiant in Grad umrechnet
func toDegrees(radians float64) float64 {
	return radians * 180 / math.Pi  
}

// MakeGeoFence ist eine Methode des GeoCodingController, die eine geographische Begrenzung (geofence) erstellt
func (gc *GeoCodingController) MakeGeoFence(lon, lat, r float64) (float64, float64, float64, float64) {
	centerLat, centerLon := lat, lon  
	radius := r  
	return getBoundingBox(centerLat, centerLon, radius) 
}

// getBoundingBox ist eine Funktion, die das Begrenzungsrechteck für eine gegebene Mitte und Radius berechnet
func getBoundingBox(centerLat, centerLon, radius float64) (minLat, minLon, maxLat, maxLon float64) {
	latDelta := (radius / 111) // 1 Grad Breitengrad entspricht etwa 111 km
	lonDelta := (radius / 111) / math.Cos(toRadians(centerLat))  // Anpassung des Längengrades an die Breitengradkrümmung

	minLat = centerLat - latDelta  // Untere Breitengradgrenze
	maxLat = centerLat + latDelta  // Obere Breitengradgrenze
	minLon = centerLon - lonDelta  // Untere Längengradgrenze
	maxLon = centerLon + lonDelta  // Obere Längengradgrenze

	return 
}

// Haversine ist eine Methode des GeoCodingController, die die Entfernung zwischen zwei geographischen Punkten berechnet
func (gc *GeoCodingController) Haversine(lat1, lon1, lat2, lon2 float64) float64 {
	lat1Rad := toRadians(lat1)  // Umwandlung des ersten Breitengrades in Radiant
	lon1Rad := toRadians(lon1)  // Umwandlung des ersten Längengrades in Radiant
	lat2Rad := toRadians(lat2)  // Umwandlung des zweiten Breitengrades in Radiant
	lon2Rad := toRadians(lon2)  // Umwandlung des zweiten Längengrades in Radiant

	deltaLat := lat2Rad - lat1Rad  // Berechnung der Differenz der Breitengrade
	deltaLon := lon2Rad - lon1Rad  // Berechnung der Differenz der Längengrade

	// Berechnung der haversine-Formel
	a := math.Pow(math.Sin(deltaLat/2), 2) + math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Pow(math.Sin(deltaLon/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return earthRadius * c  // Rückgabe der berechneten Entfernung
}
