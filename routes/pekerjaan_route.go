// package routes

// import (
// 	"database/sql"
// 	"go-fiber/app/service"
// 	"go-fiber/middleware"

// 	"github.com/gofiber/fiber/v2"
// )

// func PekerjaanRoutes(app *fiber.App, db *sql.DB) {
// 	pekerjaan := app.Group("/pekerjaan", middleware.AuthRequired())


// 	// // --- Public / User & Admin ---
// 	// pekerjaan.Get("/", middleware.AuthRequired(), func(c *fiber.Ctx) error {
// 	// 	return service.GetAllPekerjaanService(c, db)
// 	// })

// 	// pekerjaan.Get("/:id", middleware.AuthRequired(), func(c *fiber.Ctx) error {
// 	// 	return service.GetPekerjaanByIDService(c, db)
// 	// })


//     pekerjaan.Get("/alumni/:alumni_id", middleware.AdminOnly(), func(c *fiber.Ctx) error {
//         return service.GetPekerjaanByAlumniIDService(c, db)
//     })

// 	pekerjaan.Post("/", middleware.AdminOnly(), func(c *fiber.Ctx) error {
// 		return service.CreatePekerjaanService(c, db)
// 	})

// 	pekerjaan.Put("/:id", middleware.AdminOnly(), func(c *fiber.Ctx) error {
// 		return service.UpdatePekerjaanService(c, db)
// 	})

// 	pekerjaan.Delete("/:id", func(c *fiber.Ctx) error {
//     return service.SoftDeletePekerjaanService(c, db)
// 	})

// 	// pekerjaan.Delete("/:id", middleware.AdminOnly(), func(c *fiber.Ctx) error {
// 	// 	return service.DeletePekerjaanService(c, db)
// 	// })

// 	pekerjaan.Get("/search", func(c *fiber.Ctx) error {
// 		return service.GetAllPekerjaanServiceDatatable(c, db)
// 	})

// 	pekerjaan.Get("/trash", middleware.AuthRequired(), func(c *fiber.Ctx) error {
// 		return service.GetTrashedPekerjaanService(c, db)
// 	})

// 	pekerjaan.Delete("/hard/:id", middleware.AuthRequired(), func(c *fiber.Ctx) error {
// 		return service.HardDeletePekerjaanService(c, db)
// 	})

// 	pekerjaan.Put("/restore/:id", middleware.AuthRequired(), func(c *fiber.Ctx) error {
// 		return service.RestorePekerjaanService(c, db)
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

func PekerjaanRoutes(app *fiber.App, db *mongo.Database) {
	repo := repository.NewPekerjaanRepository(db, "pekerjaan_alumni", "alumni")
	pekerjaan := app.Group("/pekerjaan", middleware.AuthRequired())


	// CREATE → hanya admin
	pekerjaan.Post("/", middleware.AdminOnly(), func(c *fiber.Ctx) error {
		return service.CreatePekerjaanService(c, repo)
	})

	// UPDATE → hanya admin
	pekerjaan.Put("/:id", middleware.AdminOnly(), func(c *fiber.Ctx) error {
		return service.UpdatePekerjaanService(c, repo)
	})

	// DELETE → hanya admin
	pekerjaan.Delete("/:id", middleware.AdminOnly(), func(c *fiber.Ctx) error {
		return service.DeletePekerjaanService(c, repo)
	})

	// GET BY USER_ID → user lihat pekerjaan alumni miliknya sendiri
	// pekerjaan.Get("/user/me", func(c *fiber.Ctx) error {
	// 	// service udah otomatis ambil userID dari context
	// 	return service.GetAllPekerjaanService(c, repo)
	// })

	// GET BY ALUMNI_ID → hanya admin
	pekerjaan.Get("/alumni/:alumni_id", middleware.AdminOnly(), func(c *fiber.Ctx) error {
		return service.GetPekerjaanByAlumniIDService(c, repo)
	})

	// GET ALL → admin lihat semua, user lihat pekerjaan alumni miliknya
	pekerjaan.Get("/all", func(c *fiber.Ctx) error {
		return service.GetAllPekerjaanService(c, repo)
	})

	// GET BY ID → cek ownership kalau bukan admin
	pekerjaan.Get("/:id", func(c *fiber.Ctx) error {
		return service.GetPekerjaanByIDService(c, repo)
	})




}
