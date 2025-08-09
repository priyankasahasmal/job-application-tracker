package main

import (
	"job-application-tracker/database"
	"job-application-tracker/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.ConnectDB()

	app := fiber.New()

	app.Static("/", "./frontend")

	app.Get("/applications", handlers.GetApplications)
	app.Post("/applications", handlers.CreateApplication)
	app.Put("/applications/:id", handlers.UpdateApplication)
	app.Delete("/applications/:id", handlers.DeleteApplication)

	log.Fatal(app.Listen(":3000"))
}
