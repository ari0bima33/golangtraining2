package comment

import (
	"final-project/domain/general"
	"final-project/infra"
	"final-project/repo"

	"github.com/sirupsen/logrus"
)

type CommentUsecase struct {
	Comment CommentDataUsecaseItf
}

func NewUsecase(repo repo.Repo, conf *general.SectionService, dbList *infra.DatabaseList, logger *logrus.Logger) CommentUsecase {
	return CommentUsecase{
		Comment: newCommentDataUsecase(repo, conf, logger, dbList),
	}
}
