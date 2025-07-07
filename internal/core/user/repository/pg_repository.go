package repository

import (
	"boilerplate/internal/core/user/models"
	"boilerplate/pkg/exception"
	"boilerplate/pkg/infra/db"
	"context"
	"database/sql"
	"errors"
)

type Repository interface {
	CreateUser(ctx context.Context, userReq models.UserRegisterRequest, createdBy string) (models.UserCreateResponse, error)
	GetAllUser(ctx context.Context) ([]models.UserListResponse, error)
	GetUserByID(ctx context.Context, id int) (models.User, error)
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
	UpdateUser(ctx context.Context, userReq models.UserUpdateRequest, updatedBy string) (models.User, error)
	DeleteUser(ctx context.Context, id int, deletedBy string) error
	GetUserByEmailAndRole(ctx context.Context, email string) (models.UserDataResponse, error)
}

type UserRepo struct {
	DBList *db.DatabaseList
}

func NewUserRepo(dbList *db.DatabaseList) UserRepo {
	return UserRepo{
		DBList: dbList,
	}
}

func (u UserRepo) CreateUser(ctx context.Context, userReq models.UserRegisterRequest, createdBy string) (models.UserCreateResponse, error) {
	var response models.UserCreateResponse
	err := u.DBList.DatabaseApp.QueryRowContext(ctx, CreateUser, userReq.Name, userReq.Email, userReq.Password, userReq.Role, createdBy).Scan(&response.Name, &response.Email)
	if err != nil {
		return response, err
	}
	return response, nil

}

func (u UserRepo) GetUserByID(ctx context.Context, id int) (models.User, error) {
	var response models.User

	err := u.DBList.DatabaseApp.QueryRowContext(ctx, GetUserByID, id).Scan(&response.ID, &response.Name, &response.Email, &response.Password, &response.Role)
	if err != nil {
		if errors.Is(err, exception.ErrNotFound) {
			return response, exception.ErrNotFound
		}
		return response, err
	}

	return response, nil
}

func (u UserRepo) UpdateUser(ctx context.Context, userReq models.UserUpdateRequest, updatedBy string) (models.User, error) {
	var response models.User

	_, err := u.DBList.DatabaseApp.ExecContext(ctx, UpdateUser, userReq.Name, userReq.Email, userReq.Password, userReq.Role, updatedBy, userReq.ID)
	if err != nil {
		return response, err
	}

	err = u.DBList.DatabaseApp.QueryRowContext(ctx, GetUserByID, userReq.ID).Scan(&response.ID, &response.Name, &response.Email, &response.Password, &response.Role)
	if err != nil {
		return response, err
	}

	return response, nil
}
func (u UserRepo) DeleteUser(ctx context.Context, id int, deletedBy string) error {
	_, err := u.DBList.DatabaseApp.ExecContext(ctx, DeleteUser, deletedBy, id)
	if err != nil {
		return err
	}
	return nil
}

func (u UserRepo) GetAllUser(ctx context.Context) ([]models.UserListResponse, error) {
	var response []models.UserListResponse

	rows, err := u.DBList.DatabaseApp.QueryContext(ctx, GetAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.UserListResponse
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Role); err != nil {
			return nil, err
		}
		response = append(response, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return response, nil
}

func (u UserRepo) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	var response models.User

	err := u.DBList.DatabaseApp.QueryRowContext(ctx, GetUserByEmail, email).Scan(&response.ID, &response.Name, &response.Email, &response.Password, &response.Role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return response, exception.ErrNotFound
		}
		return response, err
	}

	return response, nil
}

func (u UserRepo) GetUserByEmailAndRole(ctx context.Context, email string) (models.UserDataResponse, error) {
	var response models.UserDataResponse

	err := u.DBList.DatabaseApp.QueryRowContext(ctx, GetUserByEmailAndRole, email).Scan(&response.ID, &response.Name, &response.Email, &response.Role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return response, exception.ErrNotFound
		}
		return response, err
	}

	return response, nil
}
