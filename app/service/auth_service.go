// package service

// import (
// 	"database/sql"
// 	"errors"
// 	"go-fiber/app/model"
// 	"go-fiber/app/repository"
// 	"go-fiber/utils"
// 	"strconv"
// 	"strings"

// 	"github.com/gofiber/fiber/v2"
// )

// func LoginService(db *sql.DB, req model.LoginRequest) (*model.LoginResponse, error) {
// 	user, passwordHash, err := repository.FindUserByUsernameOrEmail(db, req.Username)
// 	if err != nil {
// 		return nil, errors.New("username atau password salah")
// 	}

// 	if !utils.CheckPassword(req.Password, passwordHash) {
// 		return nil, errors.New("username atau password salah")
// 	}

// 	token, err := utils.GenerateToken(*user)
// 	if err != nil {
// 		return nil, errors.New("gagal generate token")
// 	}

// 	return &model.LoginResponse{
// 		User:  *user,
// 		Token: token,
// 	}, nil
// }

// func GetUsersService(db *sql.DB) fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		// --- Ambil query params ---
// 		page, _ := strconv.Atoi(c.Query("page", "1"))
// 		limit, _ := strconv.Atoi(c.Query("limit", "10"))
// 		sortBy := c.Query("sortBy", "id")
// 		order := c.Query("order", "asc")
// 		search := c.Query("search", "")
// 		offset := (page - 1) * limit

// 		// --- Whitelist sortBy biar ga SQL injection ---
// 		sortByWhitelist := map[string]string{
// 			"id":         "id",
// 			"username":   "username",
// 			"email":      "email",
// 			"password_hash": "password_hash",
// 			"role":       "role",
// 			"created_at": "created_at",
// 		}
// 		col, ok := sortByWhitelist[sortBy]
// 		if !ok {
// 			col = "id"
// 		}

// 		// --- Validasi order ---
// 		ord := "ASC"
// 		if strings.ToLower(order) == "desc" {
// 			ord = "DESC"
// 		}

// 		// --- Query ke repo ---
// 		users, err := repository.GetUsersRepo(db, search, col, ord, limit, offset)
// 		if err != nil {
// 			return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch users"})
// 		}

// 		total, err := repository.CountUsersRepo(db, search)
// 		if err != nil {
// 			return c.Status(500).JSON(fiber.Map{"error": "Failed to count users"})
// 		}

// 		// --- Hitung total halaman ---
// 		pages := 0
// 		if total > 0 {
// 			pages = (total + limit - 1) / limit
// 		}

// 		// --- Response ---
// 		response := model.UserResponse{
// 			Data: users,
// 			Meta: model.MetaInfo{
// 				Page:   page,
// 				Limit:  limit,
// 				Total:  total,
// 				Pages:  pages,
// 				SortBy: col,
// 				Order:  ord,
// 				Search: search,
// 			},
// 		}
// 		return c.JSON(response)
// 	}
// }

// versi setelah menggunakan mongoDB

package service

import (
    "context"
    "errors"
    "go-fiber/app/model"
    "go-fiber/app/repository"
    "go-fiber/utils"
    "strconv"
    "time"

    "github.com/gofiber/fiber/v2"
)

// func LoginServiceMongo(ctx context.Context, repo *repository.UserRepository, req model.LoginRequest) (*model.LoginResponse, error) {
//     user, err := repo.FindUserByUsernameOrEmail(ctx, req.Username)
//     if err != nil {
//         return nil, errors.New("username atau password salah")
//     }

//     if !utils.CheckPassword(req.Password, user.Password) {
//         return nil, errors.New("username atau password salah")
//     }

//     token, err := utils.GenerateToken(*user)
//     if err != nil {
//         return nil, errors.New("gagal generate token")
//     }

//     return &model.LoginResponse{
//         User:  *user,
//         Token: token,
//     }, nil
// }

// func GetUsersService(repo *repository.UserRepository) fiber.Handler {
//     return func(c *fiber.Ctx) error {
//         page, _ := strconv.ParseInt(c.Query("page", "1"), 10, 64)
//         limit, _ := strconv.ParseInt(c.Query("limit", "10"), 10, 64)
//         search := c.Query("search", "")
//         offset := (page - 1) * limit

//         ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//         defer cancel()

//         users, err := repo.GetUsers(ctx, search, limit, offset)
//         if err != nil {
//             return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch users"})
//         }

//         total, err := repo.CountUsers(ctx, search)
//         if err != nil {
//             return c.Status(500).JSON(fiber.Map{"error": "Failed to count users"})
//         }

//         pages := 0
//         if total > 0 {
//             pages = int((total + limit - 1) / limit)
//         }

//         return c.JSON(fiber.Map{
//             "data":  users,
//             "meta": fiber.Map{
//                 "page":   page,
//                 "limit":  limit,
//                 "total":  total,
//                 "pages":  pages,
//                 "search": search,
//             },
//         })
//     }
// }

// LoginHandler godoc
// @Summary Login user
// @Description Melakukan login dengan username/email dan password, lalu mengembalikan token JWT
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body model.LoginRequest true "Login credentials"
// @Success 200 {object} model.LoginResponse
// @Failure 400 {object} map[string]interface{} "Request body tidak valid"
// @Failure 401 {object} map[string]interface{} "Username atau password salah"
// @Router /login [post]
func LoginHandler(repo *repository.UserRepository) fiber.Handler {
    return func(c *fiber.Ctx) error {
        var req model.LoginRequest
        if err := c.BodyParser(&req); err != nil {
            return c.Status(400).JSON(fiber.Map{
                "success": false,
                "message": "Request body tidak valid",
            })
        }

        ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
        defer cancel()

        resp, err := LoginServiceMongo(ctx, repo, req)
        if err != nil {
            return c.Status(401).JSON(fiber.Map{
                "success": false,
                "message": err.Error(),
            })
        }

        return c.JSON(fiber.Map{
            "success": true,
            "message": "Login berhasil",
            "data":    resp,
        })
    }
}

// LoginServiceMongo melakukan validasi user dan generate token JWT
func LoginServiceMongo(ctx context.Context, repo repository.IUserRepository, req model.LoginRequest) (*model.LoginResponse, error) {
    user, err := repo.FindUserByUsernameOrEmail(ctx, req.Username)
    if err != nil {
        return nil, errors.New("username atau password salah")
    }

    if !utils.CheckPassword(req.Password, user.Password) {
        return nil, errors.New("username atau password salah")
    }

    token, err := utils.GenerateToken(*user)
    if err != nil {
        return nil, errors.New("gagal generate token")
    }

    return &model.LoginResponse{
        User:  *user,
        Token: token,
    }, nil
}

// GetUsersHandler godoc
// @Summary Mendapatkan daftar user
// @Description Mengambil list user dengan pagination dan search (hanya untuk user terautentikasi)
// @Tags Users
// @Accept json
// @Produce json
// @Param page query int false "Halaman (default 1)"
// @Param limit query int false "Jumlah data per halaman (default 10)"
// @Param search query string false "Kata kunci pencarian (username/email)"
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 500 {object} map[string]interface{} "Server error"
// @Security BearerAuth
// @Router /users [get]
func GetUsersHandler(repo *repository.UserRepository) fiber.Handler {
    return func(c *fiber.Ctx) error {
        page, _ := strconv.ParseInt(c.Query("page", "1"), 10, 64)
        limit, _ := strconv.ParseInt(c.Query("limit", "10"), 10, 64)
        search := c.Query("search", "")
        offset := (page - 1) * limit

        ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
        defer cancel()

        users, err := repo.GetUsers(ctx, search, limit, offset)
        if err != nil {
            return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch users"})
        }

        total, err := repo.CountUsers(ctx, search)
        if err != nil {
            return c.Status(500).JSON(fiber.Map{"error": "Failed to count users"})
        }

        pages := 0
        if total > 0 {
            pages = int((total + limit - 1) / limit)
        }

        return c.JSON(fiber.Map{
            "data":  users,
            "meta": fiber.Map{
                "page":   page,
                "limit":  limit,
                "total":  total,
                "pages":  pages,
                "search": search,
            },
        })
    }
}
