package core

import (
	"boilerplate/config"
	"boilerplate/internal/wrapper/usecase"

	"github.com/sirupsen/logrus"
	user	"boilerplate/internal/core/user/delivery"
	role	"boilerplate/internal/core/role/delivery"
)

type CoreHandler struct {
	User	user.UserHandler
	Role	role.RoleHandler
}

func NewCoreHandler(uc usecase.Usecase, conf *config.Config, log *logrus.Logger) CoreHandler {
	return CoreHandler{
		User:	user.NewUserHandler(uc, conf, log),
		Role:	role.NewRoleHandler(uc, conf, log),
	}
}
