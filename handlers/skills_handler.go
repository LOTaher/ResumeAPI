package handlers

import (
	"database/sql"
	"encoding/json"

	"github.com/LOTaher/resumeapi/models"
	"github.com/gofiber/fiber/v2"
)

func GetSkills(db *sql.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var skills models.Skills
		query := "SELECT Languages, Frameworks, DeveloperTools, Libraries FROM Skills LIMIT 1;"
		var languages, frameworks, developerTools, libraries string

		err := db.QueryRow(query).Scan(
			&languages,
			&frameworks,
			&developerTools,
			&libraries,
		)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Could not get skills",
			})
		}

		json.Unmarshal([]byte(languages), &skills.Languages)
		json.Unmarshal([]byte(frameworks), &skills.Frameworks)
		json.Unmarshal([]byte(developerTools), &skills.DeveloperTools)
		json.Unmarshal([]byte(libraries), &skills.Libraries)

		return c.JSON(skills)
	}
}
