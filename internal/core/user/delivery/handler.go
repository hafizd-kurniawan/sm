package delivery

import (
	"boilerplate/config"
	"boilerplate/internal/core/user/models"
	"boilerplate/internal/wrapper/usecase"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/sirupsen/logrus"

	"boilerplate/pkg/exception"
)

type UserHandler struct {
	Usecase usecase.Usecase
	Conf    *config.Config
	Log     *logrus.Logger
}

func NewUserHandler(uc usecase.Usecase, conf *config.Config, logger *logrus.Logger) UserHandler {
	return UserHandler{
		Usecase: uc,
		Conf:    conf,
		Log:     logger,
	}
}

func (h UserHandler) Register(ctx *fiber.Ctx) error {
	var req models.UserRegisterRequest
	init := exception.InitException(ctx, h.Conf, h.Log)

	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	resultUser, err := h.Usecase.Core.User.CreateUser(ctx.Context(), req)
	if err != nil {
		errMessage := fmt.Sprintf("Error register user: %s", err.Error())
		return exception.CreateResponse(init, fiber.StatusInternalServerError, errMessage, "", nil)
	}

	succesMessage := "User registered successfully"
	return exception.CreateResponse(init, fiber.StatusOK, succesMessage, "", resultUser)
}

func (h UserHandler) GetUserByID(ctx *fiber.Ctx) error {
	init := exception.InitException(ctx, h.Conf, h.Log)

	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		errMessage := fmt.Sprintf("Error get user by id: %s", err.Error())
		return exception.CreateResponse(init, fiber.StatusInternalServerError, errMessage, "", nil)
	}

	resultUser, err := h.Usecase.Core.User.GetUserByID(ctx.Context(), id)
	if err != nil {
		errMessage := fmt.Sprintf("Error get user by id: %s", err.Error())
		return exception.CreateResponse(init, fiber.StatusInternalServerError, errMessage, "", nil)
	}

	succesMessage := "User retrieved successfully"
	return exception.CreateResponse(init, fiber.StatusOK, succesMessage, "", resultUser)
}

func (h UserHandler) GetUserByEmail(ctx *fiber.Ctx) error {
	init := exception.InitException(ctx, h.Conf, h.Log)

	email := ctx.Params("email")

	resultUser, err := h.Usecase.Core.User.GetUserByEmail(ctx.Context(), email)
	if err != nil {
		errMessage := fmt.Sprintf("Error get user by email: %s", err.Error())
		return exception.CreateResponse(init, fiber.StatusInternalServerError, errMessage, "", nil)
	}

	succesMessage := "User retrieved successfully"
	return exception.CreateResponse(init, fiber.StatusOK, succesMessage, "", resultUser)
}

func (h UserHandler) UpdateUser(ctx *fiber.Ctx) error {
	var req models.UserUpdateRequest
	init := exception.InitException(ctx, h.Conf, h.Log)

	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	resultUser, err := h.Usecase.Core.User.UpdateUser(ctx.Context(), req)
	if err != nil {
		errMessage := fmt.Sprintf("Error update user: %s", err.Error())
		return exception.CreateResponse(init, fiber.StatusInternalServerError, errMessage, "", nil)
	}

	succesMessage := "User updated successfully"
	return exception.CreateResponse(init, fiber.StatusOK, succesMessage, "", resultUser)
}

func (h UserHandler) DeleteUser(ctx *fiber.Ctx) error {
	init := exception.InitException(ctx, h.Conf, h.Log)

	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		errMessage := fmt.Sprintf("Error delete user by id: %s", err.Error())
		return exception.CreateResponse(init, fiber.StatusInternalServerError, errMessage, "", nil)
	}

	err = h.Usecase.Core.User.DeleteUser(ctx.Context(), id)
	if err != nil {
		errMessage := fmt.Sprintf("Error delete user by id: %s", err.Error())
		return exception.CreateResponse(init, fiber.StatusInternalServerError, errMessage, "", nil)
	}

	succesMessage := "User deleted successfully"
	return exception.CreateResponse(init, fiber.StatusOK, succesMessage, "", nil)
}

func (h UserHandler) GetAllUser(ctx *fiber.Ctx) error {
	init := exception.InitException(ctx, h.Conf, h.Log)

	resultUsers, err := h.Usecase.Core.User.GetAllUser(ctx.Context())
	if err != nil {
		errMessage := fmt.Sprintf("Error get all user: %s", err.Error())
		return exception.CreateResponse(init, fiber.StatusInternalServerError, errMessage, "", nil)
	}

	succesMessage := "Users retrieved successfully"
	return exception.CreateResponse(init, fiber.StatusOK, succesMessage, "", resultUsers)
}

func (h UserHandler) Login(ctx *fiber.Ctx) error {
	var req models.UserLoginRequest
	init := exception.InitException(ctx, h.Conf, h.Log)

	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	resultUser, err := h.Usecase.Core.User.Login(ctx.Context(), req)
	if err != nil {
		errMessage := fmt.Sprintf("Error login user: %s", err.Error())
		return exception.CreateResponse(init, fiber.StatusInternalServerError, errMessage, "", nil)
	}

	succesMessage := "User login successfully"
	return exception.CreateResponse(init, fiber.StatusOK, succesMessage, "", resultUser)
}

