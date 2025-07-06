package repository

import (
	"boilerplate/internal/core/role/models"
	"boilerplate/pkg/infra/db"
	context "context"
	"fmt"
)

type Repository interface {
	CreateRole(ctx context.Context, roleReq models.RoleCreateRequest) (models.RoleCreateRequest, error)
	GetAllRole(ctx context.Context) ([]models.RoleListResponse, error)
	GetRoleByID(ctx context.Context, id int) (models.Role, error)
	GetRoleByRole(ctx context.Context, role string) (models.Role, error)
	UpdateRole(ctx context.Context, role models.RoleUpdateRequest) (models.Role, error)
	DeleteRole(ctx context.Context, id int) error
}

type RoleRepo struct {
	DBList *db.DatabaseList
}

func NewRoleRepo(dbList *db.DatabaseList) RoleRepo {
	return RoleRepo{
		DBList: dbList,
	}
}

func (r RoleRepo) CreateRole(ctx context.Context, roleReq models.RoleCreateRequest) (models.RoleCreateRequest, error) {
	var response models.RoleCreateRequest

	err := r.DBList.DatabaseApp.QueryRowContext(ctx, CreateRole, roleReq.Role).Scan(&response.Role)
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
	if err != nil {
		return response, err
	}

	return response, nil
}

func (r RoleRepo) GetRoleByRole(ctx context.Context, role string) (models.Role, error) {
	var response models.Role

	err := r.DBList.DatabaseApp.QueryRowContext(ctx, GetRoleByRole, role).Scan(&response.ID, &response.Role)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (r RoleRepo) UpdateRole(ctx context.Context, role models.RoleUpdateRequest) (models.Role, error) {
	var response models.Role

	_, err := r.DBList.DatabaseApp.ExecContext(ctx, UpdateRole, role.Role, role.ID)
	if err != nil {
		return response, err
	}

	err = r.DBList.DatabaseApp.QueryRowContext(ctx, GetRoleByID, role.ID).Scan(&response.ID, &response.Role)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (r RoleRepo) DeleteRole(ctx context.Context, id int) error {
	res, err := r.DBList.DatabaseApp.ExecContext(ctx, DeleteRole, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no role found to delete with ID %d", id)
	}

	return nil
}
