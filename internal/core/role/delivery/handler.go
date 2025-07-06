package delivery

import (
	"boilerplate/config"
	"boilerplate/internal/core/role/models"
	"boilerplate/internal/wrapper/usecase"
	"boilerplate/pkg/exception"
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

func (h RoleHandler) CreateRole(ctx *fiber.Ctx) error {
	var req models.RoleCreateRequest
	init := exception.InitException(ctx, h.Conf, h.Log)

	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	resultRole, err := h.Usecase.Core.Role.CreateRole(ctx.Context(), req)
	if err != nil {
		errMessage := fmt.Sprintf("Error create role: %s", err.Error())
		return exception.CreateResponse(init, fiber.StatusInternalServerError, errMessage, "", nil)
	}

	succesMessage := "Role created successfully"
	return exception.CreateResponse(init, fiber.StatusOK, succesMessage, "", resultRole)
}

func (h RoleHandler) GetRoleByID(ctx *fiber.Ctx) error {
	init := exception.InitException(ctx, h.Conf, h.Log)

	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		errMessage := fmt.Sprintf("Error get role by id: %s", err.Error())
		return exception.CreateResponse(init, fiber.StatusInternalServerError, errMessage, "", nil)
	}

	resultRole, err := h.Usecase.Core.Role.GetRoleByID(ctx.Context(), id)
	if err != nil {
		errMessage := fmt.Sprintf("Error get role by id: %s", err.Error())
		return exception.CreateResponse(init, fiber.StatusInternalServerError, errMessage, "", nil)
	}

	succesMessage := "Role retrieved successfully"
	return exception.CreateResponse(init, fiber.StatusOK, succesMessage, "", resultRole)
}

func (h RoleHandler) UpdateRole(ctx *fiber.Ctx) error {
	var req models.RoleUpdateRequest
	init := exception.InitException(ctx, h.Conf, h.Log)

	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	resultRole, err := h.Usecase.Core.Role.UpdateRole(ctx.Context(), req)
	if err != nil {
		errMessage := fmt.Sprintf("Error update role: %s", err.Error())
		return exception.CreateResponse(init, fiber.StatusInternalServerError, errMessage, "", nil)
	}

	succesMessage := "Role updated successfully"
	return exception.CreateResponse(init, fiber.StatusOK, succesMessage, "", resultRole)
}

func (h RoleHandler) DeleteRole(ctx *fiber.Ctx) error {
	init := exception.InitException(ctx, h.Conf, h.Log)

	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		errMessage := fmt.Sprintf("Error delete role by id: %s", err.Error())
		return exception.CreateResponse(init, fiber.StatusInternalServerError, errMessage, "", nil)
	}

	err = h.Usecase.Core.Role.DeleteRole(ctx.Context(), id)
	if err != nil {
		errMessage := fmt.Sprintf("Error delete role by id: %s", err.Error())
		return exception.CreateResponse(init, fiber.StatusInternalServerError, errMessage, "", nil)
	}

	succesMessage := "Role deleted successfully"
	return exception.CreateResponse(init, fiber.StatusOK, succesMessage, "", nil)
}
