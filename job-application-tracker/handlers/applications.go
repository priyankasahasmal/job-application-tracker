package handlers

import (
	"job-application-tracker/database"
	"job-application-tracker/models"

	"github.com/gofiber/fiber/v2"
)

func GetApplications(c *fiber.Ctx) error {
	rows, err := database.DB.Query("SELECT id, company, position, date, link, notes, status FROM applications")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()

	var apps []models.Application
	for rows.Next() {
		var app models.Application
		if err := rows.Scan(&app.ID, &app.Company, &app.Position, &app.Date, &app.Link, &app.Notes, &app.Status); err != nil {
			return c.Status(500).SendString(err.Error())
		}
		apps = append(apps, app)
	}

	return c.JSON(apps)
}

func CreateApplication(c *fiber.Ctx) error {
	app := new(models.Application)
	if err := c.BodyParser(app); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	_, err := database.DB.Exec(
		"INSERT INTO applications (company, position, date, link, notes, status) VALUES ($1, $2, $3, $4, $5, $6)",
		app.Company, app.Position, app.Date, app.Link, app.Notes, app.Status,
	)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(fiber.Map{"message": "Application added successfully"})
}

func UpdateApplication(c *fiber.Ctx) error {
	id := c.Params("id")
	app := new(models.Application)
	if err := c.BodyParser(app); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	_, err := database.DB.Exec(
		"UPDATE applications SET company=$1, position=$2, date=$3, link=$4, notes=$5, status=$6 WHERE id=$7",
		app.Company, app.Position, app.Date, app.Link, app.Notes, app.Status, id,
	)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(fiber.Map{"message": "Application updated successfully"})
}

func DeleteApplication(c *fiber.Ctx) error {
	id := c.Params("id")

	_, err := database.DB.Exec("DELETE FROM applications WHERE id=$1", id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(fiber.Map{"message": "Application deleted successfully"})
}
