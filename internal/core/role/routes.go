package role

import (
	"boilerplate/internal/middleware"
	"boilerplate/internal/wrapper/handler"
	"boilerplate/pkg/constants/role"

	"github.com/gofiber/fiber/v2"
)

func NewRoutes(api fiber.Router, handler handler.Handler) {
	// Hanya pengguna dengan peran "admin" yang dapat mengelola role.
	api.Post("/role", middleware.RoleAuthMiddleware(role.Admin), handler.Core.Role.CreateRole)
	api.Get("/role", middleware.RoleAuthMiddleware(role.Admin), handler.Core.Role.GetAllRole)
	api.Get("/role/:id", middleware.RoleAuthMiddleware(role.Admin), handler.Core.Role.GetRoleByID)
	api.Put("/role/:id", middleware.RoleAuthMiddleware(role.Admin), handler.Core.Role.UpdateRole)
	api.Delete("/role/:id", middleware.RoleAuthMiddleware(role.Admin), handler.Core.Role.DeleteRole)
}
