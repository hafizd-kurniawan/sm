package role

import (
	"boilerplate/internal/wrapper/handler"

	"github.com/gofiber/fiber/v2"
)

func NewRoutes(api fiber.Router, handler handler.Handler) {
	api.Post("/role", handler.Core.Role.CreateRole)
	api.Get("/role", handler.Core.Role.GetAllRole)
	api.Get("/role/:id", handler.Core.Role.GetRoleByID)
	api.Put("/role/:id", handler.Core.Role.UpdateRole)
	api.Delete("/role/:id", handler.Core.Role.DeleteRole)
}
