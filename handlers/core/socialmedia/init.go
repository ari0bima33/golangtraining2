package socialmedia

import (
	"final-project/domain/general"
	"final-project/usecase"

	"github.com/sirupsen/logrus"
)

type SocialmediaHandler struct {
	Socialmedia SocialmediaDataHandler
}

func NewHandler(uc usecase.Usecase, conf *general.SectionService, logger *logrus.Logger) SocialmediaHandler {
	return SocialmediaHandler{
		Socialmedia: newSocialmediaHandler(uc, conf, logger),
	}
}
