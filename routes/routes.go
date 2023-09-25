package routes

import (
	"database/sql"

	"github.com/LOTaher/resumeapi/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, db *sql.DB) {

	app.Get("/api/info", handlers.GetInfo(db))
	app.Get("/api/education", handlers.GetEducation(db))
	app.Get("/api/projects", handlers.GetProjects(db))
	app.Get("/api/experience", handlers.GetExperience(db))
	app.Get("/api/skills", handlers.GetSkills(db))

}