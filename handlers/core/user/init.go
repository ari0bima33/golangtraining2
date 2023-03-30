package user

import (
	"final-project/domain/general"
	"final-project/usecase"

	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	User UserDataHandler
}

func NewHandler(uc usecase.Usecase, conf *general.SectionService, logger *logrus.Logger) UserHandler {
	return UserHandler{
		User: newUserHandler(uc, conf, logger),
	}
}
