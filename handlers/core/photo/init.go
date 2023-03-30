package photo

import (
	"final-project/domain/general"
	"final-project/usecase"

	"github.com/sirupsen/logrus"
)

type PhotoHandler struct {
	Photo PhotoDataHandler
}

func NewHandler(uc usecase.Usecase, conf *general.SectionService, logger *logrus.Logger) PhotoHandler {
	return PhotoHandler{
		Photo: newPhotoHandler(uc, conf, logger),
	}
}
