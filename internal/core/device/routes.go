package device

import (
	"boilerplate/internal/middleware"
	"boilerplate/internal/wrapper/handler"
	"boilerplate/pkg/constants/role"

	"github.com/gofiber/fiber/v2"
)

func NewRoutes(api fiber.Router, handler handler.Handler) {
	api.Get("/devices", middleware.RoleAuthMiddleware(role.Viewer, role.Technician, role.Admin), handler.Core.Device.GetAllDevice)
	api.Get("/devices/:id", middleware.RoleAuthMiddleware(role.Viewer, role.Technician, role.Admin), handler.Core.Device.GetDeviceByID)

	api.Post("/devices", middleware.RoleAuthMiddleware(role.Technician, role.Admin), handler.Core.Device.CreateDevice)
	api.Put("/devices/:id", middleware.RoleAuthMiddleware(role.Technician, role.Admin), handler.Core.Device.UpdateDevice)
	api.Delete("/devices/:id", middleware.RoleAuthMiddleware(role.Admin), handler.Core.Device.DeleteDevice)

}

