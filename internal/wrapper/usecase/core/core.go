package core

import (
	"boilerplate/config"
	"boilerplate/internal/wrapper/repository"
	"boilerplate/pkg/infra/db"

	role "boilerplate/internal/core/role/usecase"
	user "boilerplate/internal/core/user/usecase"

	"github.com/sirupsen/logrus"
)

type CoreUsecase struct {
	User user.Usecase
	Role role.Usecase
}

func NewCoreUsecase(repo repository.Repository, conf *config.Config, dbList *db.DatabaseList, log *logrus.Logger) CoreUsecase {
	return CoreUsecase{
		User: user.NewUserUsecase(repo, conf, dbList, log),
		Role: role.NewRoleUsecase(repo, conf, dbList, log),
	}
}
