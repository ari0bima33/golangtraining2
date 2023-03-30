package socialmedia

import (
	"final-project/domain/general"
	"final-project/infra"
	"final-project/repo"

	"github.com/sirupsen/logrus"
)

type SocialmediaUsecase struct {
	Socialmedia SocialmediaDataUsecaseItf
}

func NewUsecase(repo repo.Repo, conf *general.SectionService, dbList *infra.DatabaseList, logger *logrus.Logger) SocialmediaUsecase {
	return SocialmediaUsecase{
		Socialmedia: newSocialmediaDataUsecase(repo, conf, logger, dbList),
	}
}
