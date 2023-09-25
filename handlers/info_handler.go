package handlers

import (
	"database/sql"

	"github.com/LOTaher/resumeapi/database"
	"github.com/LOTaher/resumeapi/models"
	"github.com/gofiber/fiber/v2"
)

func GetInfo(db *sql.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
	db, err := database.InitDatabase()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not initialize database",
		})
	}

	var info models.Info
	query := "SELECT * FROM Information LIMIT 1;"
	err = db.QueryRow(query).Scan(
		&info.Name,
		&info.Email,
		&info.Phone,
		&info.Website,
		&info.Availability,
		&info.Linkedin,
		&info.Github,
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not get info",
		})
	}
	
	return c.JSON(info)
}
}