package socialmedia

import (
	"final-project/infra"

	"github.com/sirupsen/logrus"
)

type SocialmediaRepo struct {
	Socialmedia SocialmediaDataRepoItf
}

func NewMasterRepo(db *infra.DatabaseList, logger *logrus.Logger) SocialmediaRepo {
	return SocialmediaRepo{
		Socialmedia: newSocialmediaDataRepo(db),
	}
}
