package handlers

import (
	"database/sql"
	"encoding/json"

	"github.com/LOTaher/resumeapi/database"
	"github.com/LOTaher/resumeapi/models"
	"github.com/gofiber/fiber/v2"
)

func GetExperience(db *sql.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
	db, err := database.InitDatabase()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not initialize database",
		})
	}

	var experiences []models.Experience
	rows, err := db.Query("SELECT Company, Position, StartDate, EndDate, Responsibilities FROM Experience")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not get experience",
		})
	}
	defer rows.Close()

	for rows.Next() {
		var experience models.Experience
		var responsibilities string

		if err := rows.Scan(
			&experience.Company,
			&experience.Position,
			&experience.StartDate,
			&experience.EndDate,
			&responsibilities,
		); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Could not scan row",
			})
		}

		json.Unmarshal([]byte(responsibilities), &experience.Responsibilities)

		experiences = append(experiences, experience)
	}

	return c.JSON(experiences)
}
}