package photo

import (
	"final-project/domain/general"
	"final-project/infra"
	"final-project/repo"

	"github.com/sirupsen/logrus"
)

type PhotoUsecase struct {
	Photo PhotoDataUsecaseItf
}

func NewUsecase(repo repo.Repo, conf *general.SectionService, dbList *infra.DatabaseList, logger *logrus.Logger) PhotoUsecase {
	return PhotoUsecase{
		Photo: newPhotoDataUsecase(repo, conf, logger, dbList),
	}
}
