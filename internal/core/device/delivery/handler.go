package delivery

import (
	"boilerplate/config"
	"boilerplate/internal/wrapper/usecase"
	"boilerplate/pkg/exception"

	"boilerplate/internal/core/device/models"
	"boilerplate/pkg/validator"
	"errors"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type DeviceHandler struct {
	Usecase usecase.Usecase
	Conf    *config.Config
	Log     *logrus.Logger
}

func NewDeviceHandler(uc usecase.Usecase, conf *config.Config, logger *logrus.Logger) DeviceHandler {
	return DeviceHandler{
		Usecase: uc,
		Conf:    conf,
		Log:     logger,
	}
}

func (h DeviceHandler) CreateDevice(ctx *fiber.Ctx) error {
	var req models.DeviceCreateRequest
	init := exception.InitException(ctx, h.Conf, h.Log)

	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	errMessage, errMessageInd := validator.ValidateDataRequest(req)
	if errMessage != "" || errMessageInd != "" {
		return exception.CreateResponse(init, fiber.StatusBadRequest, errMessage, errMessageInd, nil)
	}

	userEmail := ctx.Locals("employee_name").(string)
	resultDevice, err := h.Usecase.Core.Device.CreateDevice(ctx.Context(), req, userEmail)
	if err != nil {
		errMessage := fmt.Sprintf("Error create device: %s", err.Error())
		return exception.CreateResponse(init, fiber.StatusInternalServerError, errMessage, "", nil)
	}

	succesMessage := "Device created successfully"
	return exception.CreateResponse(init, fiber.StatusOK, succesMessage, "", resultDevice)
}

func (h DeviceHandler) GetAllDevice(ctx *fiber.Ctx) error {
	init := exception.InitException(ctx, h.Conf, h.Log)

	resultDevice, err := h.Usecase.Core.Device.GetAllDevices(ctx.Context())
	if err != nil {
		errMessage := fmt.Sprintf("Error get all device: %s", err.Error())
		return exception.CreateResponse(init, fiber.StatusInternalServerError, errMessage, "", nil)
	}

	succesMessage := "Devices retrieved successfully"
	return exception.CreateResponse(init, fiber.StatusOK, succesMessage, "", resultDevice)
}

func (h DeviceHandler) GetDeviceByID(ctx *fiber.Ctx) error {
	init := exception.InitException(ctx, h.Conf, h.Log)

	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		errMessage := fmt.Sprintf("Error get device by id: %s", err.Error())
		return exception.CreateResponse(init, fiber.StatusInternalServerError, errMessage, "", nil)
	}

	resultDevice, err := h.Usecase.Core.Device.GetDeviceByID(ctx.Context(), id)
	if err != nil {
		if errors.Is(err, exception.ErrNotFound) {
			errMessage := fmt.Sprintf("Error get device by id: %s", err.Error())
			return exception.CreateResponse(init, fiber.StatusNotFound, errMessage, "", nil)
		}

		errMessage := fmt.Sprintf("Error get device by id: %s", err.Error())
		return exception.CreateResponse(init, fiber.StatusInternalServerError, errMessage, "", nil)
	}

	succesMessage := "Device retrieved successfully"
	return exception.CreateResponse(init, fiber.StatusOK, succesMessage, "", resultDevice)
}

func (h DeviceHandler) UpdateDevice(ctx *fiber.Ctx) error {
	var req models.DeviceUpdateRequest
	init := exception.InitException(ctx, h.Conf, h.Log)

	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	errMessage, errMessageInd := validator.ValidateDataRequest(req)
	if errMessage != "" || errMessageInd != "" {
		return exception.CreateResponse(init, fiber.StatusBadRequest, errMessage, errMessageInd, nil)
	}

	userEmail := ctx.Locals("employee_name").(string)
	resultDevice, err := h.Usecase.Core.Device.UpdateDevice(ctx.Context(), req, userEmail)
	if err != nil {
		errMessage := fmt.Sprintf("Error update device: %s", err.Error())
		return exception.CreateResponse(init, fiber.StatusInternalServerError, errMessage, "", nil)
	}

	succesMessage := "Device updated successfully"
	return exception.CreateResponse(init, fiber.StatusOK, succesMessage, "", resultDevice)
}

func (h DeviceHandler) DeleteDevice(ctx *fiber.Ctx) error {
	init := exception.InitException(ctx, h.Conf, h.Log)

	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		errMessage := fmt.Sprintf("Error delete device by id: %s", err.Error())
		return exception.CreateResponse(init, fiber.StatusInternalServerError, errMessage, "", nil)
	}

	userEmail := ctx.Locals("employee_name").(string)
	err = h.Usecase.Core.Device.DeleteDevice(ctx.Context(), id, userEmail)
	if err != nil {
		errMessage := fmt.Sprintf("Error delete device by id: %s", err.Error())
		return exception.CreateResponse(init, fiber.StatusInternalServerError, errMessage, "", nil)
	}

	succesMessage := "Device deleted successfully"
	return exception.CreateResponse(init, fiber.StatusOK, succesMessage, "", nil)
}
