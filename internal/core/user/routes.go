package user

import (
	"boilerplate/internal/wrapper/handler"

	"github.com/gofiber/fiber/v2"
)

func NewRoutes(api fiber.Router, handler handler.Handler) {
	api.Post("/user/register", handler.Core.User.Register)
	api.Get("/user/get/:id", handler.Core.User.GetUserByID)
	api.Get("/user/get/email/:email", handler.Core.User.GetUserByEmail)
	api.Put("/user/update", handler.Core.User.UpdateUser)
	api.Delete("/user/delete/:id", handler.Core.User.DeleteUser)
	api.Get("/user/all", handler.Core.User.GetAllUser)
	api.Post("/user/login", handler.Core.User.Login)
}

