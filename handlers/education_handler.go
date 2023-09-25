package handlers

import (
	"database/sql"

	"github.com/LOTaher/resumeapi/models"
	"github.com/gofiber/fiber/v2"
)

func GetEducation(db *sql.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		var educations []models.Education
		rows, err := db.Query("SELECT Year, Degree, School FROM Education")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Could not get education",
			})
		}
		defer rows.Close()

		for rows.Next() {
			var education models.Education
			if err := rows.Scan(&education.Year, &education.Degree, &education.School); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "Could not scan row",
				})
			}
			educations = append(educations, education)
		}

		return c.JSON(educations)
	}
}
