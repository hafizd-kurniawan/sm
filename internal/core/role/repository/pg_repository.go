package repository

import (
	"boilerplate/internal/core/role/models"
	"boilerplate/pkg/exception"
	"boilerplate/pkg/infra/db"
	context "context"
	"database/sql"
	"errors"
)

type Repository interface {
	CreateRole(ctx context.Context, roleReq models.RoleCreateRequest, createdBy string) (models.RoleCreateResponse, error)
	GetAllRole(ctx context.Context) ([]models.RoleListResponse, error)
	GetRoleByID(ctx context.Context, id int) (models.Role, error)
	GetRoleByRole(ctx context.Context, role string) (models.Role, error)
	UpdateRole(ctx context.Context, role models.RoleUpdateRequest, updatedBy string) (models.Role, error)
	DeleteRole(ctx context.Context, id int, deletedBy string) error
}

type RoleRepo struct {
	DBList *db.DatabaseList
}

func NewRoleRepo(dbList *db.DatabaseList) RoleRepo {
	return RoleRepo{
		DBList: dbList,
	}
}

func (r RoleRepo) CreateRole(ctx context.Context, roleReq models.RoleCreateRequest, createdBy string) (models.RoleCreateResponse, error) {
	var response models.RoleCreateResponse
	err := r.DBList.DatabaseApp.QueryRowContext(ctx, CreateRole, roleReq.Role, createdBy).Scan(&response.Role)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (r RoleRepo) GetAllRole(ctx context.Context) ([]models.RoleListResponse, error) {
	var response []models.RoleListResponse

	rows, err := r.DBList.DatabaseApp.QueryContext(ctx, GetAllRole)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var role models.RoleListResponse
		if err := rows.Scan(&role.ID, &role.Role); err != nil {
			return nil, err
		}
		response = append(response, role)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return response, nil
}

func (r RoleRepo) GetRoleByID(ctx context.Context, id int) (models.Role, error) {
	var response models.Role

	err := r.DBList.DatabaseApp.QueryRowContext(ctx, GetRoleByID, id).Scan(&response.ID, &response.Role)
	if errors.Is(err, sql.ErrNoRows) {
		return response, exception.ErrNotFound
	}

	return response, nil
}

func (r RoleRepo) GetRoleByRole(ctx context.Context, role string) (models.Role, error) {
	var response models.Role

	err := r.DBList.DatabaseApp.QueryRowContext(ctx, GetRoleByRole, role).Scan(&response.ID, &response.Role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return response, exception.ErrNotFound
		}
		return response, err
	}

	return response, nil
}

func (r RoleRepo) UpdateRole(ctx context.Context, role models.RoleUpdateRequest, updatedBy string) (models.Role, error) {
	var response models.Role

	_, err := r.DBList.DatabaseApp.ExecContext(ctx, UpdateRole, role.Role, updatedBy, role.ID)
	if err != nil {
		return response, err
	}

	err = r.DBList.DatabaseApp.QueryRowContext(ctx, GetRoleByID, role.ID).Scan(&response.ID, &response.Role)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (r RoleRepo) DeleteRole(ctx context.Context, id int, deletedBy string) error {
	_, err := r.DBList.DatabaseApp.ExecContext(ctx, DeleteRole, deletedBy, id)
	if err != nil {
		return err
	}
	return nil
}
