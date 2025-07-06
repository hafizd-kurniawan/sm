package core

import (
	"boilerplate/config"
	"boilerplate/pkg/infra/db"

	role "boilerplate/internal/core/role/repository"
	user "boilerplate/internal/core/user/repository"

	"github.com/sirupsen/logrus"
)

type CoreRepository struct {
	User user.Repository
	Role role.Repository
}

func NewCoreRepository(conf *config.Config, dbList *db.DatabaseList, log *logrus.Logger) CoreRepository {
	return CoreRepository{
		User: user.NewUserRepo(dbList),
		Role: role.NewRoleRepo(dbList),
	}
}
