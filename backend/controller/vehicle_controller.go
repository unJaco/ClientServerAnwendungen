package controller

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/unJaco/UberClientServer/backend/model"
	enums "github.com/unJaco/UberClientServer/backend/model/enums"
	"gorm.io/gorm"
)

type VehicleController struct {
	db *gorm.DB
}

func NewVehicleController(db *gorm.DB) *VehicleController {
	return &VehicleController{db: db}
}

func (vc *VehicleController) CreateVehicle(c *fiber.Ctx) error {
	var vehicle model.Vehicle
	if err := c.BodyParser(&vehicle); err != nil {
		errMessage := "Invalid vehicle data"

		if errors.Is(err, gorm.ErrDuplicatedKey) {
			errMessage = "Duplicated key: Plate and / or DriverId are already taken"
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": errMessage,
		})
	}

	if result := vc.db.Create(&vehicle); result.Error != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error creating vehicle",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
		"vehicle": vehicle,
	})
}

func (vc *VehicleController) GetAllVehicles(c *fiber.Ctx) error {
	var vehicles []model.Vehicle
	vc.db.Find(&vehicles)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":  "success",
		"vehicles": vehicles,
	})
}

func (vc *VehicleController) GetVehicleByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var vehicle model.Vehicle
	err := vc.db.First(&vehicle, id).Error

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Vehicle not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"vehicle": vehicle,
	})
}

func (vc *VehicleController) GetVehicleByDriverID(c *fiber.Ctx) error {
	
	id := c.Params("id")
	var vehicle model.Vehicle
	err := vc.db.Where("driver_id = ?", id).First(&vehicle).Error

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Vehicle not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"vehicle": vehicle,
	})
}

func (vc *VehicleController) DeleteVehicleByID(c *fiber.Ctx) error {
	id := c.Params("id")
	result := vc.db.Delete(&model.Vehicle{}, id)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error deleting vehicle",
		})
	}

	if result.RowsAffected > 0 {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"success": "Vehicle was deleted successfully",
		})
	} else {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Vehicle does not exist",
		})
	}
}

func (vc *VehicleController) UpdateVehiclePosition(c *fiber.Ctx) error {

	id := c.Params("id")
	var location model.Location

	if err := c.BodyParser(&location); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid location data",
		})
	}
	result := vc.db.Model(&model.Vehicle{}).Where("id = ?", id).Updates(model.Vehicle{Location: model.Location{Lat: location.Lat, Lon: location.Lon}})

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error updating vehicle position",
		})
	}

	if result.RowsAffected > 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message":  "success",
			"location": fiber.Map{"lat": location.Lat, "lon": location.Lon},
		})
	} else {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Vehicle does not exist",
		})
	}

}

func (vc *VehicleController) UpdateVehicleStatus(c *fiber.Ctx) error {

	id := c.Params("id")
	var status = 0

	if err := c.BodyParser(&status); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid status",
		})
	}

	if status != 0 && status != 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid status",
		})
	}

	result := vc.db.Model(&model.Vehicle{}).Where("id = ?", id).Updates(model.Vehicle{Status: enums.VehhicleStatus(status)})

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error updating vehicle status",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"status":  status,
	})
}
