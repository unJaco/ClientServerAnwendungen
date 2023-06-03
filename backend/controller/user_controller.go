package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/unJaco/UberClientServer/backend/model"
	requestmodel "github.com/unJaco/UberClientServer/backend/model/requestModels"
	"gorm.io/gorm"
)

type UserController struct {
	db *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{db: db}
}

func (uc *UserController) Login(c *fiber.Ctx) error {

	var input requestmodel.LoginRequest
	if err := c.BodyParser(&input); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad Request",
		})
	}

	var user model.User
	result := uc.db.Where("email = ?", input.Email).First(&user)

	if result.Error != nil {
		return c.JSON(fiber.Map{
			"error": "email does not exist",
		})
	}

	if user.PwHash == input.PwHash {
		return c.JSON(fiber.Map{
			"message": "success",
			"user":    user,
		})
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "Unauthorized",
	})
}

func (uc *UserController) CreateUser(c *fiber.Ctx) error {

	var user model.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user data",
		})
	}

	if err := uc.db.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error creating user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
		"user":    user,
	})
}

func (uc *UserController) GetAllUsers(c *fiber.Ctx) error {
	var users []model.User
	uc.db.Find(&users)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"users":   users,
	})
}

func (uc *UserController) GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var user model.User
	err := uc.db.First(&user, id).Error

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"user":    user,
	})
}

func (uc *UserController) DeleteUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	result := uc.db.Delete(&model.User{}, id)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error deleting user",
		})
	}

	if result.RowsAffected > 0 {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"success": "User was deleted successfully",
		})
	} else {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User does not exist",
		})
	}
}
