package user

import (
	"final-project/infra"

	"github.com/sirupsen/logrus"
)

type UserRepo struct {
	User UserDataRepoItf
}

func NewMasterRepo(db *infra.DatabaseList, logger *logrus.Logger) UserRepo {
	return UserRepo{
		User: newUserDataRepo(db),
	}
}
