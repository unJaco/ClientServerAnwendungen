package controller

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/unJaco/UberClientServer/backend/model"
	enums "github.com/unJaco/UberClientServer/backend/model/enums"
	requestmodel "github.com/unJaco/UberClientServer/backend/model/requestModels"
	"gorm.io/gorm"
)

// RideController hat eine DB connection, einen GeoCodingController und einen VehicleController
type RideController struct {
	db *gorm.DB
	gc GeoCodingController
	vc VehicleController
}

// erstellen eines neuen RideControllers
func NewRideController(db *gorm.DB, gc GeoCodingController) *RideController {
	return &RideController{db: db, gc: gc}
}

func (rc *RideController) RequestRide(c *fiber.Ctx) error {

	var driveRequest requestmodel.DriveRequest

	// mitgelifernte JSON-Daten in DriveRequest umwandeln
	// wenn error dann return bad request, da JSON nicht im erwarteten Format
	if err := c.BodyParser(&driveRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid drive request data",
		})
	}

	// nutzer mit mitgeschickter CustomerId suchen
	result := rc.db.First(&model.User{}, driveRequest.CustomerId)

	// wenn error, also nicht gefunden, dann return bad request
	if result.Error != nil {
		fmt.Println("invalid id")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid customer id",
		})
	}

	fmt.Println("search")
	// nähesten fahrer von startLocation finden
	locationData, message, err := CalculateDriver(driveRequest.StartLocation, rc)

	fmt.Println(locationData)
	// wenn error, dann return internalServerError
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error creating the ride",
		})
	}

	// wenn message, also kein Fahrer gefunden, dann return StatusNotFound mit der fehler meldung
	if message != nil {

		fmt.Println("no driver")
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message" : message,
		})
	} else {
		// wenn fahrer gefunden, dann initialisieren der benötigten variablen

		driverId := uint(locationData["driverId"])
		distanceToDriver := locationData["distance"]

		startLat, startLon := driveRequest.StartLocation.Lat, driveRequest.StartLocation.Lon
		endLat, endLon := driveRequest.EndLocation.Lat, driveRequest.EndLocation.Lon

		// berechnung der distanz
		rideDistance := rc.gc.Haversine(startLat, startLon, endLat, endLon)

		var vehicle model.Vehicle

		now := time.Now()

		timeToStart := now.Add(time.Minute * time.Duration(distanceToDriver/50*60))
		arrivalTime := timeToStart.Add(time.Hour * time.Duration(rideDistance/50*60))

		// vehicle welches driverId hat suchen und in  "vehicle" speichern
		rc.db.First(&vehicle, driverId)

		/* TODO: implement a call to vehicle controller here */
		// status der vehicles auf driving setzen
		result := rc.db.Model(&model.Vehicle{}).Where("id = ?", vehicle.ID).Updates(model.Vehicle{Status: enums.VehhicleStatus(2)})

		// wenn error beim updaten des status, dann return internal server error
		if result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error updating vehicle status",
			})
		}

		// initialisieren einer Fahrt mit zugehörigen daten
		ride := model.Ride{
			VehicleId:  vehicle.ID,
			CustomerId: driveRequest.CustomerId,
			DriverId:   driverId,
			StartLon:   startLon,
			StartLat:   startLat,
			EndLon:     endLon,
			EndLat:     endLat,
			Distance:   rideDistance,
			Status:     enums.Ongoing,
			Price:      float32(rideDistance * 2),
		}

		// erstellen der Fahrt in der DB
		if err := rc.db.Create(&ride).Error; err != nil {
			/* TODO: implement a call to vehicle controller here */

			// wenn error beim erstellen, dann innerhalb der Loop
			// wenn error, dann setze vehicle status auf active und return internal server error
			rc.db.Model(&model.Vehicle{}).Where("id = ?", vehicle.ID).Updates(model.Vehicle{Status: enums.VehhicleStatus(1)})

			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error creating ride",
			})
		}

		// wenn erfolgreich, dann Map mit ride, vehicle, etc. an frontend schicken
		return c.JSON(fiber.Map{
			"ride":             ride,
			"vehicle":          vehicle,
			"distanceToDriver": distanceToDriver,
			"startTime":        timeToStart,
			"arrivalTime":      arrivalTime, // 50kmh needs to be more detailed

		})
	}

}

// berechnen des nähesten Fahrers
func CalculateDriver(location model.Location, rc *RideController) (map[string]float64, map[string]string, error) {

	// minimale und maximale Latitude und Longitude berechnen
	minLat, minLon, maxLat, maxLon := rc.gc.MakeGeoFence(location.Lon, location.Lat, 5.0)

	fmt.Println(minLat, minLon, maxLat, maxLon)
	var vehicles []model.Vehicle

	// alle vehicle innerhalb des geoFence in den array "vehicles" speichern
	/* TODO: rewrite SQL String*/
	result := rc.db.Where("lat > ? AND lat < ? AND lon > ? AND lon < ? AND status = ?", minLat, maxLat, minLon, maxLon, 1).Find(&vehicles)

	// wenn error dann dieses zurückgeben
	if result.Error != nil {
		return nil, nil, result.Error
	}

	var locationData map[string]float64

	// für jedes vehicle innerhalb des geoFence Distanz zum startPunkt berechnen
	for _, vehicle := range vehicles {
		loc := vehicle.Location

		// berechen der distanz zum startPunkt
		distance := rc.gc.Haversine(loc.Lat, loc.Lon, location.Lat, location.Lon)

		// wenn distanz geringer ist als die vorherige geringste distanz dann locationData updaten
		// wenn nicht geringer, dann nichts tun
		if distance < float64(locationData["distance"]) || locationData["distance"] == 0 {
			locationData = map[string]float64{
				"driverId": float64(vehicle.ID),
				"distance": distance,
				"lat":      loc.Lat,
				"lon":      loc.Lon,
			}
		}
	}

	// wenn mindestens ein vehicle gefunden wurde, dann return locationData
	if len(vehicles) > 0 {
		return locationData, nil, nil
	} else {
		// wenn kein vehicle gefunden, dann return map mit error message
		return nil, map[string]string{"error": "No vehicle in your area available"}, nil
	}

}

func (rc *RideController) CompleteRide(c *fiber.Ctx) error {

	// id aus der url speichern
	id := c.Params("id")

	var ride model.Ride

	// checken ob fahrt mit der id vorhanden
	err := rc.db.First(&ride, id).Error

	// wenn err, also nicht vorhanden, dann return status not found
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Ride not found",
		})
	}

	// update ride status zu complete
	result := rc.db.Model(&model.Ride{}).Where("id = ?", id).Updates(model.Ride{Status: enums.RideStatus(1)})

	// wenn error beim update, dann return internal server error
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error completing ride",
		})
	}

	/* TODO: implement a call to vehicle controller here */
	// update vehicle status zu active
	rc.db.Model(&model.Vehicle{}).Where("id = ?", ride.VehicleId).Updates(model.Vehicle{Status: enums.VehhicleStatus(1)})

	// return statusOk
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
	})

}

// fahrt bewerten
func (rc *RideController) RateRide(c *fiber.Ctx) error {

	// speichern von id und rating aus url
	id := c.Params("id")
	rating := c.Params("rating")

	// rating in integer umwandeln
	ratingInt, idErr := strconv.Atoi(rating)

	// wenn error beim umwandeln oder int <= 0 oder int >= 6 (also ungültig, da nur 1-5 erlaubt), dann return bad request
	if idErr != nil || ratingInt <= 0 || ratingInt >= 6 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid rating",
		})
	}

	var ride model.Ride

	// checken ob fahrt mit id existiert
	err := rc.db.First(&ride, id).Error

	// wenn err, also fahrt existiert nicht, dann return status not found
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Ride not found",
		})
	}

	// update rating zu mitgeschickter int, bei fahrt mit mitgeschickter id
	result := rc.db.Model(&model.Ride{}).Where("id = ?", id).Updates(model.Ride{Rating: uint(ratingInt)})

	// wenn error bei update, return internal server error
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error updating rating",
		})
	}

	// return statusOK
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"rating":  ratingInt,
	})
}

func (rc *RideController) GetAllRides(c *fiber.Ctx) error {

	var rides []model.Ride

	// speicher alle fahrten in "rides"
	rc.db.Find(&rides)

	// return StatusOK und alla fahrten
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"rides":   rides,
	})
}

// Get Ride by id
// @Description Get Ride by id.
// @Summary Get Ride by id
// @Tags Ride
// @Param id query integer true "id to filter by"
// @Produce json
// @Success 200 {object} model.Ride
// @Failure 404 {string} string "ride not found"
// @Router /api/ride/:id [get]
func (rc *RideController) GetRideByID(c *fiber.Ctx) error {

	// speicher id
	id := c.Params("id")
	var ride model.Ride

	// suche fahrt mit id und speicher in "ride"
	err := rc.db.First(&ride, id).Error

	// wenn keine fahrt mit id, dann return status not found
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Ride not found",
		})
	}

	// wenn fahrt mit id, return statusOK und fahrt
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"ride":    ride,
	})
}

// Get All Rides from one Customer
// @Description Get all Rides from one customer.
// @Summary Get all Rides from one customer
// @Tags Ride
// @Param id query integer true "id to filter by"
// @Produce json
// @Success 200 {array} model.Ride
// @Router /api/ride/customer/:id [get]
func (rc *RideController) GetAllRidesFromCustomer(c *fiber.Ctx) error {

	id := c.Params("id")
	var rides []model.Ride

	// speicher alle fahrten in "rides"
	rc.db.Where("customer_id", id).Find(&rides)

	// return StatusOK und alle fahrten
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"rides":   rides,
	})
}

func (rc *RideController) GetAllRidesFromDriver(c *fiber.Ctx) error {

	id := c.Params("id")
	var rides []model.Ride

	// speicher alle fahrten in "rides"
	rc.db.Where("driver_id", id).Find(&rides)

	// return StatusOK und alle fahrten
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"rides":   rides,
	})
}



