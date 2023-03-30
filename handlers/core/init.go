package core

import (
	"final-project/domain/general"
	"final-project/handlers/core/authorization"
	"final-project/handlers/core/comment"
	"final-project/handlers/core/photo"
	"final-project/handlers/core/socialmedia"
	"final-project/handlers/core/user"
	"final-project/usecase"

	"github.com/sirupsen/logrus"
)

type Handler struct {
	Token       authorization.TokenHandler
	Public      authorization.PublicHandler
	User        user.UserHandler
	Photo       photo.PhotoHandler
	Comment     comment.CommentHandler
	Socialmedia socialmedia.SocialmediaHandler
}

func NewHandler(uc usecase.Usecase, conf *general.SectionService, logger *logrus.Logger) Handler {
	return Handler{
		Token:       authorization.NewTokenHandler(conf, logger),
		Public:      authorization.NewPublicHandler(conf, logger),
		User:        user.NewHandler(uc, conf, logger),
		Photo:       photo.NewHandler(uc, conf, logger),
		Comment:     comment.NewHandler(uc, conf, logger),
		Socialmedia: socialmedia.NewHandler(uc, conf, logger),
	}
}
