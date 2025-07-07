package delivery

import (
	"boilerplate/config"
	"boilerplate/internal/core/user/models"
	"boilerplate/internal/wrapper/usecase"
	"errors"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/sirupsen/logrus"

	"boilerplate/pkg/exception"
	"boilerplate/pkg/validator"
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

	errMessage, errMessageInd := validator.ValidateDataRequest(req)
	if errMessage != "" || errMessageInd != "" {
		return exception.CreateResponse(init, fiber.StatusBadRequest, errMessage, errMessageInd, nil)
	}

	userEmail := ""
	resultUser, err := h.Usecase.Core.User.CreateUser(ctx.Context(), req, userEmail)
	if err != nil {
		errMessage = fmt.Sprintf("Error register user: %s", err.Error())
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
		if errors.Is(err, exception.ErrNotFound) {
			errMessage := fmt.Sprintf("Error get user by id: %s", err.Error())
			return exception.CreateResponse(init, fiber.StatusNotFound, errMessage, "", nil)
		}

		errMessage := fmt.Sprintf("Error get user by id: %s", err.Error())
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

	errMessage, errMessageInd := validator.ValidateDataRequest(req)
	if errMessage != "" || errMessageInd != "" {
		return exception.CreateResponse(init, fiber.StatusBadRequest, errMessage, errMessageInd, nil)
	}

	userEmail := ctx.Locals("employee_name").(string)
	resultUser, err := h.Usecase.Core.User.UpdateUser(ctx.Context(), req, userEmail)
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

	userEmail := ctx.Locals("employee_name").(string)
	err = h.Usecase.Core.User.DeleteUser(ctx.Context(), id, userEmail)
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

	errMessage, errMessageInd := validator.ValidateDataRequest(req)
	if errMessage != "" || errMessageInd != "" {
		return exception.CreateResponse(init, fiber.StatusBadRequest, errMessage, errMessageInd, nil)
	}

	resultUser, err := h.Usecase.Core.User.Login(ctx.Context(), req)
	if err != nil {
		errMessage := fmt.Sprintf("Error login user: %s", err.Error())
		return exception.CreateResponse(init, fiber.StatusInternalServerError, errMessage, "", nil)
	}

	succesMessage := "User login successfully"
	return exception.CreateResponse(init, fiber.StatusOK, succesMessage, "", resultUser)
}

func (h UserHandler) GetMe(ctx *fiber.Ctx) error {
	init := exception.InitException(ctx, h.Conf, h.Log)

	employeeName := ctx.Locals("employee_name")

	if employeeName == "" {
		errMessage := fmt.Sprintf("Error get user by email: %s", exception.ErrUnauthorized)
		return exception.CreateResponse(init, fiber.StatusUnauthorized, errMessage, "", nil)
	}

	email := employeeName.(string)
	resultUser, err := h.Usecase.Core.User.GetUserByEmailAndRole(ctx.Context(), email)
	if err != nil {
		errMessage := fmt.Sprintf("Error get user by email: %s", err.Error())
		return exception.CreateResponse(init, fiber.StatusInternalServerError, errMessage, "", nil)
	}

	response := models.UserDataResponse{
		ID:    resultUser.ID,
		Email: resultUser.Email,
		Name:  resultUser.Name,
		Role:  resultUser.Role,
	}

	succesMessage := "User retrieved successfully"
	return exception.CreateResponse(init, fiber.StatusOK, succesMessage, "", response)
}
