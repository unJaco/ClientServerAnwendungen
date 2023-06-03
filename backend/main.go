package main

import (
	"fmt"
	"log"
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
	"github.com/unJaco/UberClientServer/backend/controller"
	_ "github.com/unJaco/UberClientServer/backend/docs"
	"github.com/unJaco/UberClientServer/backend/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// @title UberClientServer API
// @version 1.0
// @description This is an API to manage Users, Vehicles and Rides.
// @host localhost:8080
// @BasePath /
func main() {

	// laden der .env file
	errEnv := godotenv.Load("../.env")

	// wenn error beim laden, dann beenden des Programms
	if errEnv != nil {
		log.Fatalf(errEnv.Error())
		log.Fatalf("Error loading .env file")
	}

	// lesen und speichern der variablen aus der .env file
	user_name := os.Getenv("user_name")
	password := os.Getenv("password")
	address := os.Getenv("address")
	db_name := os.Getenv("db_name")

	// öffnen der DB Connection mit den variablen
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user_name, password, address, db_name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// bei error programm beenden
	if err != nil {
		log.Println(err)
		panic("failed to connect database")

	} else {
		log.Println("Connection Established")
	}

	// automatisches migrieren der Models in die Datenbank
	// für jedes model wird eine eigene Tablle erstellt
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Ride{})
	db.AutoMigrate(&model.Vehicle{})

	// initialisieren der benötigten controller
	userController := controller.NewUserController(db)
	vehicleController := controller.NewVehicleController(db)
	rideController := controller.NewRideController(db, *&controller.GeoCodingController{})

	// erstllen einer neuen fiber app
	app := fiber.New()

	// CORS konfiguration setzen um anfragen des frontends (localhost:3000) zu erlauben
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowCredentials: true,
	}))

	// initialisren der routen
	setupRoutes(app, userController, vehicleController, rideController)

	// app auf port 8080 starten
	app.Listen(":8080")

}

func setupRoutes(app *fiber.App, uc *controller.UserController, vc *controller.VehicleController, rc *controller.RideController) {

	// hinzufügen des swagger endpoints
	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "/swagger/doc.json",
		DeepLinking: false,
	}))

	// hinzuügen der endpoints und der groups
	api := app.Group("/api")

	user := api.Group("/user")

	// beispiel : localhost:8080/api/user/all -> im userController (uc) wird GetAllUsers aufgerufen
	user.Get("/all", uc.GetAllUsers)
	user.Get("/getUser/:id", uc.GetUserByID)

	user.Post("/create", uc.CreateUser)
	user.Delete("/:id", uc.DeleteUserByID)

	// TODO: why is this post
	user.Post("/login", uc.Login)

	vehicle := api.Group("/vehicle")

	vehicle.Get("/all", vc.GetAllVehicles)
	vehicle.Get("/:id", vc.GetVehicleByID)
	vehicle.Get("/driverID/:id", vc.GetVehicleByDriverID)

	vehicle.Post("/create", vc.CreateVehicle)
	vehicle.Post("/:id/position", vc.UpdateVehiclePosition)
	vehicle.Post("/:id/status", vc.UpdateVehicleStatus)

	// TODO : do we need this
	vehicle.Delete("/delete/:id", vc.DeleteVehicleByID)

	ride := api.Group("/ride")

	ride.Get("/all", rc.GetAllRides)
	ride.Get("/:id", rc.GetRideByID)
	ride.Get("customer/:id", rc.GetAllRidesFromCustomer)
	ride.Get("driver/:id", rc.GetAllRidesFromDriver)

	ride.Post("/request", rc.RequestRide)
	ride.Post("/complete/:id", rc.CompleteRide)
	ride.Post("rate/:id/:rating", rc.RateRide)

}
