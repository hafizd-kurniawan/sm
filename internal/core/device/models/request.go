package models

type DeviceCreateRequest struct {
	Name     string `json:"name" validate:"required"`
	Location string `json:"location"`
	Status   string `json:"status" validate:"required,oneof=online offline"`
}

type DeviceUpdateRequest struct {
	ID       int    `json:"id" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Location string `json:"location"`
	Status   string `json:"status" validate:"required,oneof=online offline"`
}

