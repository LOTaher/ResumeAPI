package handlers

import (
	"database/sql"

	"github.com/LOTaher/resumeapi/models"
	"github.com/gofiber/fiber/v2"
)

func GetProjects(db *sql.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		var projects []models.Projects
		query := "SELECT Name, Description, Technologies FROM Projects;"
		rows, err := db.Query(query)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Could not get projects",
			})
		}
		defer rows.Close()

		for rows.Next() {
			var project models.Projects
			if err := rows.Scan(
				&project.Name,
				&project.Description,
				&project.Technologies,
			); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "Could not scan row",
				})
			}
			
			projects = append(projects, project)
		}

		return c.JSON(projects)
	}
}
