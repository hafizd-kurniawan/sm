package usecase

import (
	"boilerplate/config"
	repo "boilerplate/internal/wrapper/repository"
	"boilerplate/pkg/exception"
	"boilerplate/pkg/infra/db"
	"context"
	"errors"

	"fmt"

	"github.com/sirupsen/logrus"

	"boilerplate/internal/core/role/models"
)

type Usecase interface {
	CreateRole(ctx context.Context, roleReq models.RoleCreateRequest, createdBy string) (models.RoleCreateResponse, error)
	GetAllRole(ctx context.Context) ([]models.RoleListResponse, error)
	GetRoleByID(ctx context.Context, id int) (models.Role, error)
	UpdateRole(ctx context.Context, role models.RoleUpdateRequest, updatedBy string) (models.Role, error)
	DeleteRole(ctx context.Context, id int, deletedBy string) error
}

type RoleUsecase struct {
	Repo   repo.Repository
	Conf   *config.Config
	DBList *db.DatabaseList
	Log    *logrus.Logger
}

func NewRoleUsecase(repository repo.Repository, conf *config.Config, dbList *db.DatabaseList, logger *logrus.Logger) RoleUsecase {
	return RoleUsecase{
		Repo:   repository,
		Conf:   conf,
		DBList: dbList,
		Log:    logger,
	}
}

func (u RoleUsecase) CreateRole(ctx context.Context, roleReq models.RoleCreateRequest, createdBy string) (models.RoleCreateResponse, error) {
	var response models.RoleCreateResponse

	_, err := u.Repo.Core.Role.GetRoleByRole(ctx, roleReq.Role)
	fmt.Println(err)
	if err == nil {
		return response, fmt.Errorf("%w: role with name '%s' already exists", exception.ErrConflict, roleReq.Role)
	}

	if !errors.Is(err, exception.ErrNotFound) {
		u.Log.Errorf("failed to check role existence: %v", err)
		return response, err
	}

	role, err := u.Repo.Core.Role.CreateRole(ctx, roleReq, createdBy)
	if err != nil {
		u.Log.Error(err)
		return role, err
	}
	return role, nil
}

func (u RoleUsecase) GetAllRole(ctx context.Context) ([]models.RoleListResponse, error) {
	role, err := u.Repo.Core.Role.GetAllRole(ctx)
	if err != nil {
		u.Log.Error(err)
		return role, err
	}
	return role, nil
}

func (u RoleUsecase) GetRoleByID(ctx context.Context, id int) (models.Role, error) {
	role, err := u.Repo.Core.Role.GetRoleByID(ctx, id)
	if err != nil {
		u.Log.Error(err)
		return models.Role{}, err
	}
	return role, nil
}

func (u RoleUsecase) UpdateRole(ctx context.Context, role models.RoleUpdateRequest, updatedBy string) (models.Role, error) {
	updatedRole, err := u.Repo.Core.Role.UpdateRole(ctx, role, updatedBy)
	if err != nil {
		u.Log.Error(err)
		return models.Role{}, err
	}
	return updatedRole, nil
}

func (u RoleUsecase) DeleteRole(ctx context.Context, id int, deletedBy string) error {
	err := u.Repo.Core.Role.DeleteRole(ctx, id, deletedBy)
	if err != nil {
		u.Log.Error(err)
		return err
	}
	return nil
}
