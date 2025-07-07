package usecase

import (
	"boilerplate/config"
	repo "boilerplate/internal/wrapper/repository"
	"boilerplate/pkg/exception"
	"boilerplate/pkg/infra/db"
	"context"
	"errors"
	"fmt"

	"boilerplate/internal/core/device/models"

	"github.com/sirupsen/logrus"
)

type Usecase interface {
	CreateDevice(ctx context.Context, req models.DeviceCreateRequest, createdBy string) (models.Device, error)
	GetAllDevices(ctx context.Context) ([]models.Device, error)
	GetDeviceByID(ctx context.Context, id int) (models.Device, error)
	UpdateDevice(ctx context.Context, req models.DeviceUpdateRequest, updatedBy string) (models.Device, error)
	DeleteDevice(ctx context.Context, id int, deletedBy string) error
}

type DeviceUsecase struct {
	Repo   repo.Repository
	Conf   *config.Config
	DBList *db.DatabaseList
	Log    *logrus.Logger
}

func NewDeviceUsecase(repository repo.Repository, conf *config.Config, dbList *db.DatabaseList, logger *logrus.Logger) DeviceUsecase {
	return DeviceUsecase{
		Repo:   repository,
		Conf:   conf,
		DBList: dbList,
		Log:    logger,
	}
}

func (u DeviceUsecase) CreateDevice(ctx context.Context, req models.DeviceCreateRequest, createdBy string) (models.Device, error) {
	_, err := u.Repo.Core.Device.GetDeviceByName(ctx, req.Name)
	if err == nil {
		u.Log.Errorf("device duplikat dengan nama '%s'", req.Name)
		return models.Device{}, exception.ErrConflict
	}

	device, err := u.Repo.Core.Device.CreateDevice(ctx, req, createdBy)
	if err != nil {
		u.Log.Errorf("failed to create device: %v", err)
		return models.Device{}, err
	}

	return device, nil
}

func (u DeviceUsecase) GetAllDevices(ctx context.Context) ([]models.Device, error) {
	devices, err := u.Repo.Core.Device.GetAllDevices(ctx)
	if err != nil {
		u.Log.Errorf("failed to get all devices: %v", err)
		return nil, err
	}
	return devices, nil
}

func (u DeviceUsecase) GetDeviceByID(ctx context.Context, id int) (models.Device, error) {
	device, err := u.Repo.Core.Device.GetDeviceByID(ctx, id)
	if err != nil {
		if !errors.Is(err, exception.ErrNotFound) {
			u.Log.Errorf("failed to get device by id: %v", err)
		}
		return models.Device{}, err
	}
	return device, nil
}

func (u DeviceUsecase) UpdateDevice(ctx context.Context, req models.DeviceUpdateRequest, updatedBy string) (models.Device, error) {
	device, err := u.Repo.Core.Device.GetDeviceByID(ctx, req.ID)
	if err != nil {
		if errors.Is(err, exception.ErrNotFound) {
			return models.Device{}, err
		}
		u.Log.Errorf("UpdateDevice: failed to get device: %v", err)
		return models.Device{}, fmt.Errorf("cannot get device: %w", err)
	}

	if _, err := u.Repo.Core.Device.UpdateDevice(ctx, req, updatedBy); err != nil {
		u.Log.Errorf("UpdateDevice: failed to update: %v", err)
		return models.Device{}, fmt.Errorf("update failed: %w", err)
	}

	device, err = u.Repo.Core.Device.GetDeviceByID(ctx, req.ID)
	if err != nil {
		u.Log.Errorf("UpdateDevice: failed to fetch updated device: %v", err)
		return models.Device{}, fmt.Errorf("fetch updated failed: %w", err)
	}

	return device, nil
}

func (u DeviceUsecase) DeleteDevice(ctx context.Context, id int, deletedBy string) error {
	err := u.Repo.Core.Device.DeleteDevice(ctx, id, deletedBy)
	if err != nil {
		if !errors.Is(err, exception.ErrNotFound) {
			u.Log.Errorf("failed to delete device: %v", err)
		}
		return err
	}
	return nil
}
