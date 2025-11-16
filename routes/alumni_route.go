// package routes

// import (
// 	"database/sql"
// 	"go-fiber/app/service"
// 	"go-fiber/middleware"

// 	"github.com/gofiber/fiber/v2"
// )

// func AlumniRoutes(app *fiber.App, db *sql.DB) {
// 	alumni := app.Group("/alumni", middleware.AuthRequired()) // semua butuh login

// 	// // GET all alumni → admin & user boleh
// 	// alumni.Get("/", func(c *fiber.Ctx) error {
// 	// 	return service.GetAllAlumniService(c, db)
// 	// })

// 	// // GET alumni by ID → admin & user boleh
// 	// alumni.Get("/:id", func(c *fiber.Ctx) error {
// 	// 	return service.GetAlumniByIDService(c, db)
// 	// })

// 	// CREATE alumni → hanya admin
// 	alumni.Post("/", middleware.AdminOnly(), func(c *fiber.Ctx) error {
// 		return service.CreateAlumniService(c, db)
// 	})

// 	// UPDATE alumni → hanya admin
// 	alumni.Put("/:id", middleware.AdminOnly(), func(c *fiber.Ctx) error {
// 		return service.UpdateAlumniService(c, db)
// 	})

// 	// DELETE alumni → hanya admin
// 	alumni.Delete("/:id", middleware.AdminOnly(), func(c *fiber.Ctx) error {
// 		return service.DeleteAlumniService(c, db)
// 	})

// 	alumni.Get("/search", func(c *fiber.Ctx) error {
// 		return service.GetAllAlumniServiceDatatable(c, db)
// 	})

// 	alumni.Get("/statistik", func(c *fiber.Ctx) error {
// 		return service.CountAlumniByStatusService(c, db)
// 	})

// }

package routes

import (
	"go-fiber/app/repository"
	"go-fiber/app/service"
	"go-fiber/middleware"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func AlumniRoutes(app *fiber.App, db *mongo.Database) {
	// inject repository ke service
	repo := repository.NewAlumniRepository(db, "alumni")
	alumni := app.Group("/alumni", middleware.AuthRequired())

	// CREATE → hanya admin
	alumni.Post("/", middleware.AdminOnly(), func(c *fiber.Ctx) error {
		return service.CreateAlumniService(c, repo)
	})

	// UPDATE → hanya admin
	alumni.Put("/:id", middleware.AdminOnly(), func(c *fiber.Ctx) error {
		return service.UpdateAlumniService(c, repo)
	})

	// DELETE → hanya admin
	alumni.Delete("/:id", middleware.AdminOnly(), func(c *fiber.Ctx) error {
		return service.DeleteAlumniService(c, repo)
	})

	// GET semua alumni → admin & user boleh
	alumni.Get("/", func(c *fiber.Ctx) error {
		return service.GetAllAlumniService(c, repo)
	})

	// GET alumni berdasarkan ID (opsional, belum dibuat servicenya)
	// alumni.Get("/:id", func(c *fiber.Ctx) error {
	// 	return service.GetAlumniByIDService(c, repo)
	// })

	// Statistik (belum kita migrasi, nanti setelah pekerjaan_alumni siap)
	// alumni.Get("/statistik", func(c *fiber.Ctx) error {
	// 	return service.CountAlumniByStatusService(c, repo)
	// })
}
