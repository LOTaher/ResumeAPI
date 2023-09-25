package handlers

import (
	"database/sql"
	"encoding/json"

	"github.com/LOTaher/resumeapi/database"
	"github.com/LOTaher/resumeapi/models"
	"github.com/gofiber/fiber/v2"
)

func GetProjects(db *sql.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
	db, err := database.InitDatabase()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not initialize database",
		})
	}

	var projects []models.Projects
	rows, err := db.Query("SELECT Name, Description, Technologies FROM Projects")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not get projects",
		})
	}
	defer rows.Close()

	for rows.Next() {
		var project models.Projects
		var description string
		var technologies string

		if err := rows.Scan(&project.Name, &description, &technologies); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Could not scan row",
			})
		}
		
		json.Unmarshal([]byte(description), &project.Description)
		json.Unmarshal([]byte(technologies), &project.Technologies)

		projects = append(projects, project)
	}

	return c.JSON(projects)
}
}