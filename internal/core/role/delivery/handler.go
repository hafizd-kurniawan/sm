package delivery

import (
	"boilerplate/config"
	"boilerplate/internal/core/role/models"
	"boilerplate/internal/wrapper/usecase"
	"boilerplate/pkg/exception"
	"boilerplate/pkg/validator"
	"errors"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type RoleHandler struct {
	Usecase usecase.Usecase
	Conf    *config.Config
	Log     *logrus.Logger
}

func NewRoleHandler(uc usecase.Usecase, conf *config.Config, logger *logrus.Logger) RoleHandler {
	return RoleHandler{
		Usecase: uc,
		Conf:    conf,
		Log:     logger,
	}
}

// GetAllRole retrieves a list of all roles.
// @Summary Get all roles (Admin Only)
// @Description Retrieves a complete list of all active roles in the system. Requires administrator access.
// @Tags Roles
// @Produce json
// @Success 200 {array} models.RoleResponse "List of roles retrieved successfully"
// @Failure 401 {object} exception.ResponseData "Unauthorized"
// @Failure 403 {object} exception.ResponseData "Forbidden"
// @Failure 500 {object} exception.ResponseData "Internal Server Error"
// @Security BearerAuth
// @Router /role [get]
func (h RoleHandler) GetAllRole(ctx *fiber.Ctx) error {
	init := exception.InitException(ctx, h.Conf, h.Log)

	resultRole, err := h.Usecase.Core.Role.GetAllRole(ctx.Context())
	if err != nil {
		errMessage := fmt.Sprintf("Error get all role: %s", err.Error())
		return exception.CreateResponse(init, fiber.StatusInternalServerError, errMessage, "", nil)
	}

	succesMessage := "Roles retrieved successfully"
	return exception.CreateResponse(init, fiber.StatusOK, succesMessage, "", resultRole)
}

// CreateRole handles the creation of a new role.
// @Summary Create a new role (Admin Only)
// @Description Creates a new role in the system. Requires administrator access.
// @Tags Roles
// @Accept json
// @Produce json
// @Param role body models.RoleCreateRequest true "Payload to create a new role"
// @Success 201 {object} models.RoleResponse "Role created successfully"
// @Failure 400 {object} exception.ResponseData "Bad Request - Invalid input data"
// @Failure 401 {object} exception.ResponseData "Unauthorized - Missing or invalid token"
// @Failure 403 {object} exception.ResponseData "Forbidden - User does not have admin privileges"
// @Failure 409 {object} exception.ResponseData "Conflict - Role with this name already exists"
// @Security BearerAuth
// @Router /role [post]
func (h RoleHandler) CreateRole(ctx *fiber.Ctx) error {
	var req models.RoleCreateRequest
	init := exception.InitException(ctx, h.Conf, h.Log)

	if err := ctx.BodyParser(&req); err != nil {
		return exception.CreateResponse(init, fiber.StatusBadRequest, "Invalid request body", "", nil)
	}

	errMessage, errMessageInd := validator.ValidateDataRequest(req)
	if errMessage != "" || errMessageInd != "" {
		return exception.CreateResponse(init, fiber.StatusBadRequest, errMessage, errMessageInd, nil)
	}

	userEmail := ctx.Locals("employee_name").(string)
	resultRole, err := h.Usecase.Core.Role.CreateRole(ctx.Context(), req, userEmail)
	if err != nil {
		errMessage = fmt.Sprintf("Error create role: %s", err.Error())
		return exception.CreateResponse(init, fiber.StatusInternalServerError, errMessage, "", nil)
	}

	succesMessage := "Role created successfully"
	return exception.CreateResponse(init, fiber.StatusOK, succesMessage, "", resultRole)
}

// GetRoleByID retrieves a single role by its ID.
// @Summary Get a role by ID (Admin Only)
// @Description Retrieves the details of a specific role by its unique ID. Requires administrator access.
// @Tags Roles
// @Produce json
// @Param id path int true "Role ID"
// @Success 200 {object} models.RoleResponse "Role details retrieved successfully"
// @Failure 400 {object} exception.ResponseData "Bad Request - Invalid ID format"
// @Failure 401 {object} exception.ResponseData "Unauthorized"
// @Failure 403 {object} exception.ResponseData "Forbidden"
// @Failure 404 {object} exception.ResponseData "Not Found - Role with the specified ID does not exist"
// @Security BearerAuth
// @Router /role/{id} [get]
func (h RoleHandler) GetRoleByID(ctx *fiber.Ctx) error {
	init := exception.InitException(ctx, h.Conf, h.Log)

	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		errMessage := fmt.Sprintf("Error get role by id: %s", err.Error())
		return exception.CreateResponse(init, fiber.StatusInternalServerError, errMessage, "", nil)
	}

	resultRole, err := h.Usecase.Core.Role.GetRoleByID(ctx.Context(), id)
	if err != nil {
		if errors.Is(err, exception.ErrNotFound) {
			errMessage := fmt.Sprintf("Error get role by id: %s", err.Error())
			return exception.CreateResponse(init, fiber.StatusNotFound, errMessage, "", nil)
		}

		errMessage := fmt.Sprintf("Error get role by id: %s", err.Error())
		return exception.CreateResponse(init, fiber.StatusInternalServerError, errMessage, "", nil)
	}

	succesMessage := "Role retrieved successfully"
	return exception.CreateResponse(init, fiber.StatusOK, succesMessage, "", resultRole)
}

// UpdateRole updates an existing role.
// @Summary Update a role (Admin Only)
// @Description Updates the name of an existing role. Requires administrator access.
// @Tags Roles
// @Accept json
// @Produce json
// @Param id path int true "Role ID to update"
// @Param role body models.RoleUpdateRequest true "Payload to update the role"
// @Success 200 {object} models.RoleResponse "Role updated successfully"
// @Failure 400 {object} exception.ResponseData "Bad Request - Invalid input or ID"
// @Failure 401 {object} exception.ResponseData "Unauthorized"
// @Failure 403 {object} exception.ResponseData "Forbidden"
// @Failure 404 {object} exception.ResponseData "Not Found - Role with the specified ID does not exist"
// @Security BearerAuth
// @Router /role/{id} [put]
func (h RoleHandler) UpdateRole(ctx *fiber.Ctx) error {
	var req models.RoleUpdateRequest
	init := exception.InitException(ctx, h.Conf, h.Log)

	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	errMessage, errMessageInd := validator.ValidateDataRequest(req)
	if errMessage != "" || errMessageInd != "" {
		return exception.CreateResponse(init, fiber.StatusBadRequest, errMessage, errMessageInd, nil)
	}

	userEmail := ctx.Locals("employee_name").(string)
	resultRole, err := h.Usecase.Core.Role.UpdateRole(ctx.Context(), req, userEmail)
	if err != nil {
		errMessage = fmt.Sprintf("Error update role: %s", err.Error())
		return exception.CreateResponse(init, fiber.StatusInternalServerError, errMessage, "", nil)
	}

	succesMessage := "Role updated successfully"
	return exception.CreateResponse(init, fiber.StatusOK, succesMessage, "", resultRole)
}

// DeleteRole soft-deletes a role.
// @Summary Soft delete a role (Admin Only)
// @Description Marks a role as deleted. The role is not permanently removed from the database. Requires administrator access.
// @Tags Roles
// @Produce json
// @Param id path int true "Role ID to delete"
// @Success 200 {object} object "Success message"
// @Failure 400 {object} exception.ResponseData "Bad Request - Invalid ID format"
// @Failure 401 {object} exception.ResponseData "Unauthorized"
// @Failure 403 {object} exception.ResponseData "Forbidden"
// @Failure 404 {object} exception.ResponseData "Not Found - Role with the specified ID does not exist"
// @Security BearerAuth
// @Router /role/{id} [delete]
func (h RoleHandler) DeleteRole(ctx *fiber.Ctx) error {
	init := exception.InitException(ctx, h.Conf, h.Log)

	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		errMessage := fmt.Sprintf("Error delete role by id: %s", err.Error())
		return exception.CreateResponse(init, fiber.StatusInternalServerError, errMessage, "", nil)
	}

	userEmail := ctx.Locals("employee_name").(string)
	err = h.Usecase.Core.Role.DeleteRole(ctx.Context(), id, userEmail)
	if err != nil {
		errMessage := fmt.Sprintf("Error delete role by id: %s", err.Error())
		return exception.CreateResponse(init, fiber.StatusInternalServerError, errMessage, "", nil)
	}

	succesMessage := "Role deleted successfully"
	return exception.CreateResponse(init, fiber.StatusOK, succesMessage, "", nil)
}
