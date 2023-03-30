package usecase

import (
	"final-project/domain/general"
	"final-project/infra"
	"final-project/repo"
	"final-project/usecase/comment"
	"final-project/usecase/photo"
	"final-project/usecase/socialmedia"
	"final-project/usecase/user"

	"github.com/sirupsen/logrus"
)

type Usecase struct {
	User        user.UserUsecase
	Photo       photo.PhotoUsecase
	Comment     comment.CommentUsecase
	Socialmedia socialmedia.SocialmediaUsecase
}

func NewUsecase(repo repo.Repo, conf *general.SectionService, dbList *infra.DatabaseList, logger *logrus.Logger) Usecase {
	return Usecase{
		User:        user.NewUsecase(repo, conf, dbList, logger),
		Photo:       photo.NewUsecase(repo, conf, dbList, logger),
		Comment:     comment.NewUsecase(repo, conf, dbList, logger),
		Socialmedia: socialmedia.NewUsecase(repo, conf, dbList, logger),
	}
}
