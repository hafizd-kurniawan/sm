package repository

import (
	"boilerplate/internal/core/user/models"
	"boilerplate/pkg/infra/db"
	"context"
	"fmt"
)

type Repository interface {
	CreateUser(ctx context.Context, userReq models.UserRegisterRequest) (models.UserCreateResponse, error)
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
	GetUserByID(ctx context.Context, id int) (models.User, error)
	UpdateUser(ctx context.Context, userReq models.UserUpdateRequest) (models.User, error)
	DeleteUser(ctx context.Context, id int) error
	GetAllUser(ctx context.Context) ([]models.UserListResponse, error)
	Login(ctx context.Context, userReq models.UserLoginRequest) (models.User, error)
}

type UserRepo struct {
	DBList *db.DatabaseList
}

func NewUserRepo(dbList *db.DatabaseList) UserRepo {
	return UserRepo{
		DBList: dbList,
	}
}

func (u UserRepo) CreateUser(ctx context.Context, userReq models.UserRegisterRequest) (models.UserCreateResponse, error) {
	var response models.UserCreateResponse
	err := u.DBList.DatabaseApp.QueryRowContext(ctx, CreateUser, userReq.Name, userReq.Email, userReq.Password, userReq.Role).Scan(&response.Name, &response.Email)
	if err != nil {
		return response, err
	}
	return response, nil

}
func (u UserRepo) GetUserByID(ctx context.Context, id int) (models.User, error) {
	var response models.User

	err := u.DBList.DatabaseApp.QueryRowContext(ctx, GetUserByID, id).Scan(&response.ID, &response.Name, &response.Email, &response.Password, &response.Role)
	if err != nil {
		return response, err
	}

	return response, nil
}
func (u UserRepo) UpdateUser(ctx context.Context, userReq models.UserUpdateRequest) (models.User, error) {
	var response models.User

	_, err := u.DBList.DatabaseApp.ExecContext(ctx, UpdateUser, userReq.Name, userReq.Email, userReq.Password, userReq.Role, userReq.ID)
	if err != nil {
		return response, err
	}

	err = u.DBList.DatabaseApp.QueryRowContext(ctx, GetUserByID, userReq.ID).Scan(&response.ID, &response.Name, &response.Email, &response.Password, &response.Role)
	if err != nil {
		return response, err
	}

	return response, nil
}
func (u UserRepo) DeleteUser(ctx context.Context, id int) error {
	res, err := u.DBList.DatabaseApp.ExecContext(ctx, DeleteUser, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no user found to delete with ID %d", id)
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
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role); err != nil {
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
		return response, err
	}

	return response, nil
}

func (u UserRepo) Login(ctx context.Context, userReq models.UserLoginRequest) (models.User, error) {
	var response models.User

	err := u.DBList.DatabaseApp.QueryRowContext(ctx, Login, userReq.Email, userReq.Password).Scan(&response.ID, &response.Name, &response.Email, &response.Password, &response.Role)
	if err != nil {
		return response, err
	}

	return response, nil

}
