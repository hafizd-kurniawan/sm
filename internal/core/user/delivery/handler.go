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

// Register handles new user registration.
// @Summary Register a new user
// @Description Creates a new user account. This is a public endpoint.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param user body models.UserRegisterRequest true "User Registration Payload"
// @Success 200 {object} models.UserCreateResponse
// @Failure 400 {object} exception.ResponseData "Bad Request"
// @Failure 409 {object} exception.ResponseData "Conflict - Email already exists"
// @Failure 500 {object} exception.ResponseData "Internal Server Error"
// @Router /user/register [post]
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

// GetUserByID retrieves a single user by their ID.
// @Summary Get a user by ID (Admin only)
// @Description Retrieves details of a specific user by their ID. This endpoint is restricted to users with the 'admin' role.
// @Tags Users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.User "Successfully retrieved user"
// @Failure 400 {object} exception.ResponseData "Bad Request - Invalid ID format"
// @Failure 401 {object} exception.ResponseData "Unauthorized"
// @Failure 403 {object} exception.ResponseData "Forbidden"
// @Failure 404 {object} exception.ResponseData "Not Found - User not found"
// @Security BearerAuth
// @Router /user/get/{id} [get]
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

// UpdateUser updates a user's information.
// @Summary Update a user (Admin only)
// @Description Updates a user's details, such as name, email, and role. This endpoint is restricted to users with the 'admin' role.
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID to update"
// @Param user body models.UserUpdateRequest true "User Update Payload"
// @Success 200 {object} models.User "Successfully updated user"
// @Failure 400 {object} exception.ResponseData "Bad Request - Invalid input or ID"
// @Failure 401 {object} exception.ResponseData "Unauthorized"
// @Failure 403 {object} exception.ResponseData "Forbidden"
// @Failure 404 {object} exception.ResponseData "Not Found - User not found"
// @Security BearerAuth
// @Router /user/update/{id} [put]
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

// DeleteUser soft-deletes a user.
// @Summary Soft delete a user (Admin only)
// @Description Marks a user as deleted in the system. The user is not permanently removed. This endpoint is restricted to users with the 'admin' role.
// @Tags Users
// @Produce json
// @Param id path int true "User ID to delete"
// @Success 200 {object} object "Success message"
// @Failure 400 {object} exception.ResponseData "Bad Request - Invalid ID format"
// @Failure 401 {object} exception.ResponseData "Unauthorized"
// @Failure 403 {object} exception.ResponseData "Forbidden"
// @Failure 404 {object} exception.ResponseData "Not Found - User not found"
// @Security BearerAuth
// @Router /user/delete/{id} [delete]
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

// GetAllUser retrieves a list of all users.
// @Summary Get all users (Admin only)
// @Description Retrieves a list of all users. This endpoint is restricted to users with the 'admin' role.
// @Tags Users
// @Produce json
// @Success 200 {array} models.UserListResponse "Successfully retrieved list of users"
// @Failure 401 {object} exception.ResponseData "Unauthorized - Invalid or missing token"
// @Failure 403 {object} exception.ResponseData "Forbidden - User does not have admin role"
// @Security BearerAuth
// @Router /user/all [get]
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

// Login handles user login.
// @Summary User Login
// @Description Authenticate a user and receive a JWT token. This is a public endpoint.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param credentials body models.UserLoginRequest true "User Credentials"
// @Success 200 {object} models.LoginResponse "Successfully authenticated"
// @Failure 400 {object} exception.ResponseData "Bad Request - Invalid input"
// @Failure 401 {object} exception.ResponseData "Unauthorized - Invalid credentials"
// @Failure 404 {object} exception.ResponseData "Not Found - User not found"
// @Router /user/login [post]
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

// GetMe retrieves the profile of the currently authenticated user.
// @Summary Get my profile
// @Description Retrieves the profile details for the user associated with the token. Accessible by all authenticated roles.
// @Tags Users
// @Produce json
// @Success 200 {object} models.UserDataResponse "Successfully retrieved profile"
// @Failure 401 {object} exception.ResponseData "Unauthorized - Invalid or missing token"
// @Failure 404 {object} exception.ResponseData "Not Found - User from token not found"
// @Security BearerAuth
// @Router /me [get]
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
