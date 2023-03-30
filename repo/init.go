package repo

import (
	"final-project/infra"
	"final-project/repo/comment"
	"final-project/repo/photo"
	"final-project/repo/socialmedia"
	"final-project/repo/user"

	"github.com/sirupsen/logrus"
)

type Repo struct {
	User       user.UserRepo
	Photo      photo.PhotoRepo
	Comment    comment.CommentRepo
	Sociamedia socialmedia.SocialmediaRepo
}

func NewRepo(db *infra.DatabaseList, logger *logrus.Logger) Repo {
	return Repo{
		User:       user.NewMasterRepo(db, logger),
		Photo:      photo.NewMasterRepo(db, logger),
		Comment:    comment.NewMasterRepo(db, logger),
		Sociamedia: socialmedia.NewMasterRepo(db, logger),
	}
}
