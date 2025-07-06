package usecase

import (
	"boilerplate/config"
	"boilerplate/internal/core/user/models"
	repo "boilerplate/internal/wrapper/repository"
	"boilerplate/pkg/infra/db"
	"boilerplate/pkg/utils"
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
)

type Usecase interface {
	CreateUser(ctx context.Context, userReq models.UserRegisterRequest) (models.UserCreateResponse, error)
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
	GetUserByID(ctx context.Context, id int) (models.User, error)
	UpdateUser(ctx context.Context, userReq models.UserUpdateRequest) (models.User, error)
	DeleteUser(ctx context.Context, id int) error
	GetAllUser(ctx context.Context) ([]models.UserListResponse, error)
	Login(ctx context.Context, userReq models.UserLoginRequest) (models.LoginResponse, error)
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

func (u UserUsecase) CreateUser(ctx context.Context, userReq models.UserRegisterRequest) (models.UserCreateResponse, error) {
	// cek apakah user sudah ada
	user, err := u.Repo.Core.User.GetUserByEmail(ctx, userReq.Email)
	if user.ID != 0 {
		return models.UserCreateResponse{}, fmt.Errorf("email sudah ada")
	}

	hashPassword, err := utils.HashingPassword(userReq.Password)
	if err != nil {
		return models.UserCreateResponse{}, err
	}
	createUser := models.UserRegisterRequest{
		Email:    userReq.Email,
		Name:     userReq.Name,
		Password: hashPassword,
		Role:     userReq.Role,
	}

	_, err = u.Repo.Core.User.CreateUser(ctx, createUser)
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
		return user, err
	}
	return user, nil
}

func (u UserUsecase) GetUserByID(ctx context.Context, id int) (models.User, error) {
	user, err := u.Repo.Core.User.GetUserByID(ctx, id)
	if err != nil {
		u.Log.Error(err)
		return user, err
	}
	return user, nil
}

func (u UserUsecase) UpdateUser(ctx context.Context, userReq models.UserUpdateRequest) (models.User, error) {
	user, err := u.Repo.Core.User.UpdateUser(ctx, userReq)
	if err != nil {
		u.Log.Error(err)
		return user, err
	}
	return user, nil
}

func (u UserUsecase) DeleteUser(ctx context.Context, id int) error {
	err := u.Repo.Core.User.DeleteUser(ctx, id)

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
	if user.ID == 0 {
		return response, fmt.Errorf("user not found")
	}

	if isMatchPassword := utils.CheckHashedPassword(user.Password, userReq.Password); !isMatchPassword {
		return response, fmt.Errorf("password is not match")
	}

	userReq.Password = user.Password
	user, err = u.Repo.Core.User.Login(ctx, userReq)
	if err != nil {
		u.Log.Error(err)
		return response, err
	}

	userRole, err := u.Repo.Core.Role.GetRoleByID(ctx, user.Role)
	if err != nil {
		u.Log.Error(err)
		return response, err
	}

	token, err := utils.GenereateJWT(u.Conf, user.Email, userRole.Role)

	return models.LoginResponse{
		UsesrID:  user.ID,
		Username: user.Name,
		Password: user.Password,
		Email:    user.Email,
		Token:    token,
	}, nil
}
