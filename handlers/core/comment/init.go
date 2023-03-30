package comment

import (
	"final-project/domain/general"
	"final-project/usecase"

	"github.com/sirupsen/logrus"
)

type CommentHandler struct {
	Comment CommentDataHandler
}

func NewHandler(uc usecase.Usecase, conf *general.SectionService, logger *logrus.Logger) CommentHandler {
	return CommentHandler{
		Comment: newCommentHandler(uc, conf, logger),
	}
}
