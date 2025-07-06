package usecase

import (
	"boilerplate/config"
	repo "boilerplate/internal/wrapper/repository"
	"boilerplate/pkg/infra/db"
	"context"

	"github.com/sirupsen/logrus"

	"boilerplate/internal/core/role/models"
)

type Usecase interface {
	CreateRole(ctx context.Context, roleReq models.RoleCreateRequest) (models.RoleCreateRequest, error)
	GetAllRole(ctx context.Context) ([]models.RoleListResponse, error)
	GetRoleByID(ctx context.Context, id int) (models.Role, error)
	UpdateRole(ctx context.Context, role models.RoleUpdateRequest) (models.Role, error)
	DeleteRole(ctx context.Context, id int) error
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

func (u RoleUsecase) CreateRole(ctx context.Context, roleReq models.RoleCreateRequest) (models.RoleCreateRequest, error) {
	role, err := u.Repo.Core.Role.CreateRole(ctx, roleReq)
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

func (u RoleUsecase) UpdateRole(ctx context.Context, role models.RoleUpdateRequest) (models.Role, error) {
	updatedRole, err := u.Repo.Core.Role.UpdateRole(ctx, role)
	if err != nil {
		u.Log.Error(err)
		return models.Role{}, err
	}
	return updatedRole, nil
}

func (u RoleUsecase) DeleteRole(ctx context.Context, id int) error {
	err := u.Repo.Core.Role.DeleteRole(ctx, id)
	if err != nil {
		u.Log.Error(err)
		return err
	}
	return nil
}
