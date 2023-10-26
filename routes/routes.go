package routes

import (
	"database/sql"

	"github.com/LOTaher/resumeapi/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, db *sql.DB) {

	app.Get("/api", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to my Resume API! 👋")
	})

	app.Get("/api/info", handlers.GetInfo(db))
	app.Get("/api/education", handlers.GetEducation(db))

	app.Get("/api/projects", handlers.GetProjects(db))
    // app.Get("/api/projects/:id", handlers.GetProjectsById(db))

	app.Get("/api/experience", handlers.GetExperience(db))
    // app.Get("/api/experience/:id", handlers.GetExperienceById(db))

	app.Get("/api/skills", handlers.GetSkills(db))
    app.Get("/api/skills/languages", handlers.GetSkillsByLanguages(db))
    app.Get("/api/skills/frameworks", handlers.GetSkillsByFrameworks(db))
    app.Get("/api/skills/tools", handlers.GetSkillsByTools(db))
    app.Get("/api/skills/libraries", handlers.GetSkillsByLibraries(db))

}
