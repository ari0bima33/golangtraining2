package comment

import (
	"final-project/infra"

	"github.com/sirupsen/logrus"
)

type CommentRepo struct {
	Comment CommentDataRepoItf
}

func NewMasterRepo(db *infra.DatabaseList, logger *logrus.Logger) CommentRepo {
	return CommentRepo{
		Comment: newCommentDataRepo(db),
	}
}
