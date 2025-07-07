package user

import (
	"boilerplate/internal/middleware"
	"boilerplate/internal/wrapper/handler"
	"boilerplate/pkg/constants/role"

	"github.com/gofiber/fiber/v2"
)

func NewRoutes(api fiber.Router, handler handler.Handler) {
	api.Post("/user/register", handler.Core.User.Register)
	api.Post("/user/login", handler.Core.User.Login)
	api.Get("/me", middleware.RoleAuthMiddleware(role.Admin, role.Technician, role.Viewer), handler.Core.User.GetMe)

	api.Get("/user/all", middleware.RoleAuthMiddleware(role.Admin), handler.Core.User.GetAllUser)
	api.Get("/user/get/:id", middleware.RoleAuthMiddleware(role.Admin), handler.Core.User.GetUserByID)
	api.Put("/user/update/:id", middleware.RoleAuthMiddleware(role.Admin), handler.Core.User.UpdateUser)
	api.Delete("/user/delete/:id", middleware.RoleAuthMiddleware(role.Admin), handler.Core.User.DeleteUser)
}
