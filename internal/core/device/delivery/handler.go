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

// CreateDevice handles the creation of a new device.
// @Summary Create a new device
// @Description Creates a new device with the provided details. Only accessible by Technicians and Admins.
// @Tags Devices
// @Accept json
// @Produce json
// @Param device body models.DeviceCreateRequest true "Device Create Payload"
// @Success 201 {object} models.Device "Successfully created device"
// @Failure 400 {object} exception.ResponseData "Bad Request - Invalid input"
// @Failure 401 {object} exception.ResponseData "Unauthorized - Invalid or missing token"
// @Failure 403 {object} exception.ResponseData "Forbidden - User does not have the required role"
// @Failure 500 {object} exception.ResponseData "Internal Server Error"
// @Security BearerAuth
// @Router /devices [post]
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

// GetAllDevice handles the request to get all devices.
// @Summary Get all active devices
// @Description Retrieves a list of all devices that have not been soft-deleted. Accessible by all authenticated users.
// @Tags Devices
// @Produce json
// @Success 200 {array} models.Device
// @Failure 401 {object} exception.ResponseData "Unauthorized"
// @Failure 500 {object} exception.ResponseData "Internal Server Error"
// @Security BearerAuth
// @Router /devices [get]
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

// GetDeviceByID handles the request to get a device by its ID.
// @Summary Get a single device by ID
// @Description Retrieves details of a specific device. Accessible by all authenticated users.
// @Tags Devices
// @Produce json
// @Param id path int true "Device ID"
// @Success 200 {object} models.Device
// @Failure 400 {object} exception.ResponseData "Invalid ID format"
// @Failure 401 {object} exception.ResponseData "Unauthorized"
// @Failure 404 {object} exception.ResponseData "Device not found"
// @Security BearerAuth
// @Router /devices/{id} [get]
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

// UpdateDevice handles the update of a device's data.
// @Summary Update an existing device
// @Description Updates a device's details. Only accessible by Technicians and Admins.
// @Tags Devices
// @Accept json
// @Produce json
// @Param id path int true "Device ID"
// @Param device body models.DeviceUpdateRequest true "Device Update Payload"
// @Success 200 {object} models.Device "Successfully updated device"
// @Failure 400 {object} exception.ResponseData "Bad Request - Invalid input"
// @Failure 401 {object} exception.ResponseData "Unauthorized"
// @Failure 403 {object} exception.ResponseData "Forbidden"
// @Failure 404 {object} exception.ResponseData "Device not found"
// @Security BearerAuth
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

// DeleteDevice handles the soft deletion of a device.
// @Summary Soft delete a device
// @Description Marks a device as deleted. Only accessible by Admins.
// @Tags Devices
// @Produce json
// @Param id path int true "Device ID"
// @Success 200 {object} object "Success message"
// @Failure 400 {object} exception.ResponseData "Invalid ID format"
// @Failure 401 {object} exception.ResponseData "Unauthorized"
// @Failure 403 {object} exception.ResponseData "Forbidden"
// @Failure 404 {object} exception.ResponseData "Device not found"
// @Security BearerAuth
// @Router /devices/{id} [delete]
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
