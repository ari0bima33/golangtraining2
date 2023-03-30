package user

import (
	"final-project/domain/general"
	"final-project/infra"
	"final-project/repo"

	"github.com/sirupsen/logrus"
)

type UserUsecase struct {
	User UserDataUsecaseItf
}

func NewUsecase(repo repo.Repo, conf *general.SectionService, dbList *infra.DatabaseList, logger *logrus.Logger) UserUsecase {
	return UserUsecase{
		User: newUserDataUsecase(repo, conf, logger, dbList),
	}
}
