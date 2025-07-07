package repository

import (
	"boilerplate/internal/core/device/models"
	"boilerplate/pkg/exception"
	"boilerplate/pkg/infra/db"
	"context"
	"database/sql"
	"errors"
)

type Repository interface {
	CreateDevice(ctx context.Context, req models.DeviceCreateRequest, createdBy string) (models.Device, error)
	GetAllDevices(ctx context.Context) ([]models.Device, error)
	GetDeviceByID(ctx context.Context, id int) (models.Device, error)
	UpdateDevice(ctx context.Context, req models.DeviceUpdateRequest, updatedBy string) (models.Device, error)
	DeleteDevice(ctx context.Context, id int, deletedBy string) error
	GetDeviceByName(ctx context.Context, name string) (models.Device, error)
}

type DeviceRepo struct {
	DBList *db.DatabaseList
}

func NewDeviceRepo(dbList *db.DatabaseList) DeviceRepo {
	return DeviceRepo{
		DBList: dbList,
	}
}

func (r DeviceRepo) CreateDevice(ctx context.Context, req models.DeviceCreateRequest, createdBy string) (models.Device, error) {
	var device models.Device
	err := r.DBList.DatabaseApp.QueryRowContext(
		ctx,
		CreateDevice,
		req.Name,
		req.Location,
		req.Status,
		createdBy,
	).Scan(
		&device.ID,
		&device.Name,
		&device.Location,
		&device.Status,
	)

	if err != nil {
		return models.Device{}, err
	}

	return device, nil
}

func (r DeviceRepo) GetAllDevices(ctx context.Context) ([]models.Device, error) {
	var devices []models.Device

	rows, err := r.DBList.DatabaseApp.QueryContext(ctx, GetAllDevices)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var device models.Device
		if err := rows.Scan(
			&device.ID,
			&device.Name,
			&device.Location,
			&device.Status,
			&device.CreatedAt,
			&device.UpdatedAt,
		); err != nil {
			return nil, err
		}
		devices = append(devices, device)
	}

	return devices, nil
}

func (r DeviceRepo) GetDeviceByID(ctx context.Context, id int) (models.Device, error) {
	var device models.Device
	err := r.DBList.DatabaseApp.QueryRowContext(ctx, GetDeviceByID, id).Scan(
		&device.ID,
		&device.Name,
		&device.Location,
		&device.Status,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Device{}, exception.ErrNotFound
		}
		return models.Device{}, err
	}

	return device, nil
}

func (r DeviceRepo) UpdateDevice(ctx context.Context, req models.DeviceUpdateRequest, updatedBy string) (models.Device, error) {
	var device models.Device
	res, err := r.DBList.DatabaseApp.ExecContext(ctx, UpdateDevice, req.Name, req.Location, req.Status, updatedBy, req.ID)
	if err != nil {
		return models.Device{}, err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		// Jika tidak ada baris yang terpengaruh, berarti device dengan ID tersebut tidak ditemukan.
		return models.Device{}, exception.ErrNotFound
	}
	return device, nil

}

func (r DeviceRepo) DeleteDevice(ctx context.Context, id int, deletedBy string) error {
	res, err := r.DBList.DatabaseApp.ExecContext(ctx, DeleteDevice, deletedBy, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return exception.ErrNotFound
	}

	return nil
}

func (r DeviceRepo) GetDeviceByName(ctx context.Context, name string) (models.Device, error) {
	var device models.Device
	err := r.DBList.DatabaseApp.QueryRowContext(ctx, GetDeviceByName, name).Scan(
		&device.ID,
		&device.Name,
		&device.Location,
		&device.Status,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Device{}, exception.ErrNotFound
		}
		return models.Device{}, err
	}

	return device, nil
}
