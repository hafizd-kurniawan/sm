package usecase

import (
	"boilerplate/config"
	"boilerplate/internal/core/user/models"
	repo "boilerplate/internal/wrapper/repository"
	"boilerplate/pkg/exception"
	"boilerplate/pkg/infra/db"
	"boilerplate/pkg/utils"
	"context"
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
)

type Usecase interface {
	CreateUser(ctx context.Context, userReq models.UserRegisterRequest, createdBy string) (models.UserCreateResponse, error)
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
	GetUserByID(ctx context.Context, id int) (models.User, error)
	UpdateUser(ctx context.Context, userReq models.UserUpdateRequest, updatedBy string) (models.User, error)
	DeleteUser(ctx context.Context, id int, deletedBy string) error
	GetAllUser(ctx context.Context) ([]models.UserListResponse, error)
	Login(ctx context.Context, userReq models.UserLoginRequest) (models.LoginResponse, error)
	GetUserByEmailAndRole(ctx context.Context, email string) (models.UserDataResponse, error)
}

type UserUsecase struct {
	Repo   repo.Repository
	Conf   *config.Config
	DBList *db.DatabaseList
	Log    *logrus.Logger
}

func NewUserUsecase(repository repo.Repository, conf *config.Config, dbList *db.DatabaseList, logger *logrus.Logger) UserUsecase {
	return UserUsecase{
		Repo:   repository,
		Conf:   conf,
		DBList: dbList,
		Log:    logger,
	}
}

func (u UserUsecase) CreateUser(ctx context.Context, userReq models.UserRegisterRequest, createdBy string) (models.UserCreateResponse, error) {
	var response models.UserCreateResponse
	_, err := u.Repo.Core.User.GetUserByEmail(ctx, userReq.Email)
	if err == nil {
		return response, fmt.Errorf("%w: email sudah ada", exception.ErrConflict)
	}

	hashPassword, err := utils.HashingPassword(userReq.Password)
	if err != nil {
		return response, err
	}

	createUser := models.UserRegisterRequest{
		Email:    userReq.Email,
		Name:     userReq.Name,
		Password: hashPassword,
		Role:     userReq.Role,
	}

	_, err = u.Repo.Core.User.CreateUser(ctx, createUser, createdBy)
	if err != nil {
		return models.UserCreateResponse{}, err
	}

	return models.UserCreateResponse{
		Name:  createUser.Name,
		Email: createUser.Email,
	}, nil
}

func (u UserUsecase) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	user, err := u.Repo.Core.User.GetUserByEmail(ctx, email)
	if err != nil {
		u.Log.Error(err)
		return user, fmt.Errorf("%w: user not found", exception.ErrNotFound)
	}
	return user, nil
}

func (u UserUsecase) GetUserByID(ctx context.Context, id int) (models.User, error) {
	user, err := u.Repo.Core.User.GetUserByID(ctx, id)
	if err != nil {
		u.Log.Error(err)
		return user, fmt.Errorf("%w: user not found", exception.ErrNotFound)
	}
	return user, nil
}

func (u UserUsecase) UpdateUser(ctx context.Context, userReq models.UserUpdateRequest, updatedBy string) (models.User, error) {

	user, err := u.Repo.Core.User.GetUserByID(ctx, userReq.ID)
	if err != nil {
		u.Log.Error(err)
		return user, err
	}

	hashPassword, err := utils.HashingPassword(userReq.Password)
	if err != nil {
		u.Log.Error(err)
		return user, err
	}
	user.Password = hashPassword

	user, err = u.Repo.Core.User.UpdateUser(ctx, userReq, updatedBy)
	if err != nil {
		u.Log.Error(err)
		return user, err
	}
	return user, nil
}

func (u UserUsecase) DeleteUser(ctx context.Context, id int, deletedBy string) error {
	err := u.Repo.Core.User.DeleteUser(ctx, id, deletedBy)

	if err != nil {
		u.Log.Error(err)
		return err
	}
	return nil
}

func (u UserUsecase) GetAllUser(ctx context.Context) ([]models.UserListResponse, error) {
	users, err := u.Repo.Core.User.GetAllUser(ctx)
	if err != nil {
		u.Log.Error(err)
		return nil, err
	}
	return users, nil
}

func (u UserUsecase) Login(ctx context.Context, userReq models.UserLoginRequest) (models.LoginResponse, error) {
	var response models.LoginResponse

	user, err := u.Repo.Core.User.GetUserByEmail(ctx, userReq.Email)
	if err != nil {
		u.Log.Error(err)
		return response, fmt.Errorf("%w: user not found", exception.ErrNotFound)
	}

	if isMatchPassword := utils.CheckHashedPassword(user.Password, userReq.Password); !isMatchPassword {
		u.Log.Error(err)
		return response, fmt.Errorf("%w: password is not match", exception.ErrUnauthorized)
	}

	userRole, err := u.Repo.Core.Role.GetRoleByID(ctx, user.Role)
	if err != nil {
		u.Log.Error(err)
		return response, err
	}

	token, err := utils.GenereateJWT(u.Conf, user.Email, userRole.Role)

	return models.LoginResponse{
		Username: user.Name,
		Email:    user.Email,
		Token:    token,
	}, nil
}

func (u UserUsecase) GetUserByEmailAndRole(ctx context.Context, email string) (models.UserDataResponse, error) {
	user, err := u.Repo.Core.User.GetUserByEmailAndRole(ctx, email)
	if err != nil {
		if errors.Is(err, exception.ErrNotFound) {
			return models.UserDataResponse{}, exception.ErrNotFound
		}
		u.Log.Error(err)
		return models.UserDataResponse{}, err
	}
	return user, nil
}
