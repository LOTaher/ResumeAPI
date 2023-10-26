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

func GetSkillsByLanguages(db *sql.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var skills models.Skills
		query := "SELECT Languages FROM Skills LIMIT 1;"
		var languages string

		err := db.QueryRow(query).Scan(
			&languages,
		)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Could not get languages",
			})
		}

		json.Unmarshal([]byte(languages), &skills.Languages)

		return c.JSON(fiber.Map{"Languages": skills.Languages})
	}
}

func GetSkillsByFrameworks(db *sql.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var skills models.Skills
		query := "SELECT Frameworks FROM Skills LIMIT 1;"
		var frameworks string

		err := db.QueryRow(query).Scan(
			&frameworks,
		)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Could not get frameworks",
			})
		}

		json.Unmarshal([]byte(frameworks), &skills.Frameworks)

		return c.JSON(fiber.Map{"Frameworks": skills.Frameworks})
	}
}

func GetSkillsByTools(db *sql.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var skills models.Skills
		query := "SELECT DeveloperTools FROM Skills LIMIT 1;"
		var developerTools string

		err := db.QueryRow(query).Scan(
			&developerTools,
		)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Could not get developer tools",
			})
		}

		json.Unmarshal([]byte(developerTools), &skills.DeveloperTools)

		return c.JSON(fiber.Map{"Tools": skills.DeveloperTools})
	}
}

func GetSkillsByLibraries(db *sql.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var skills models.Skills
		query := "SELECT Libraries FROM Skills LIMIT 1;"
		var libraries string

		err := db.QueryRow(query).Scan(
			&libraries,
		)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Could not get libraries",
			})
		}

		json.Unmarshal([]byte(libraries), &skills.Libraries)

		return c.JSON(fiber.Map{"Libraries": skills.Libraries})
	}
}
