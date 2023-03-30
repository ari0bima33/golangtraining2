package socialmedia

import (
	"errors"

	"final-project/domain/general"
	dc "final-project/domain/socialmedia"
	"final-project/infra"
	"final-project/repo"
	rc "final-project/repo/socialmedia"
	"final-project/utils"

	"github.com/sirupsen/logrus"
)

type SocialmediaDataUsecaseItf interface {
	GetList() ([]dc.SocialmediaList, string, error)
	CreateSocialmedia(data dc.CreateSocialmediaRequest, userID int64) (*dc.CreateSocialmediaResponse, string, error)
	UpdateSocialmedia(data dc.UpdateSocialmediaRequest, ID int64, userID int64) (*dc.UpdateSocialmediaResponse, string, error)
	DeleteSocialmedia(ID int64, userID int64) (*dc.DeleteSocialmediaResponse, string, error)
}

type SocialmediaDataUsecase struct {
	Repo   rc.SocialmediaDataRepoItf
	DBList *infra.DatabaseList
	Conf   *general.SectionService
	Log    *logrus.Logger
}

func newSocialmediaDataUsecase(r repo.Repo, conf *general.SectionService, logger *logrus.Logger, dbList *infra.DatabaseList) SocialmediaDataUsecase {
	return SocialmediaDataUsecase{
		Repo:   r.Sociamedia.Socialmedia,
		Conf:   conf,
		Log:    logger,
		DBList: dbList,
	}
}

func (uu SocialmediaDataUsecase) GetList() ([]dc.SocialmediaList, string, error) {
	var result []dc.SocialmediaList

	resultSocialmediaList, err := uu.Repo.GetList()
	if err != nil {
		uu.Log.WithField("request", utils.StructToString(nil)).WithError(err).Errorf("fail to get list socialmedia")
		return nil, "fail to get list socialmedia", err
	}

	if len(resultSocialmediaList) > 0 {
		for _, v := range resultSocialmediaList {
			user := dc.User{
				ID:              v.UserID,
				Username:        v.Username,
				ProfileImageUrl: v.ProfileImageUrl,
			}
			socialmedia := dc.SocialmediaList{
				ID:             v.ID,
				Name:           v.Name,
				SocialMediaUrl: v.SocialMediaUrl,
				UserID:         v.UserID,
				CreatedAt:      v.CreatedAt,
				UpdatedAt:      v.UpdatedAt,
				User:           user,
			}
			result = append(result, socialmedia)
		}
	}

	return result, "success get list socialmedia", nil
}

func (uu SocialmediaDataUsecase) CreateSocialmedia(data dc.CreateSocialmediaRequest, userID int64) (*dc.CreateSocialmediaResponse, string, error) {

	resultSocialmedia, err := uu.Repo.InsertSocialmedia(nil, data, userID)
	if err != nil {
		uu.Log.WithField("request", utils.StructToString(data)).WithError(err).Errorf("fail to Insert Socialmedia")
		return nil, "fail to create socialmedia", err
	}

	if resultSocialmedia == nil {
		uu.Log.WithField("request", utils.StructToString(data)).WithError(err).Errorf("fail to Insert Socialmedia")
		return nil, "fail to create socialmedia", err
	}

	return resultSocialmedia, "success create socialmedia", nil
}

func (uu SocialmediaDataUsecase) UpdateSocialmedia(data dc.UpdateSocialmediaRequest, ID int64, userID int64) (*dc.UpdateSocialmediaResponse, string, error) {

	socialmedia, err := uu.Repo.GetByIDUserID(ID, userID)
	if err != nil {
		uu.Log.WithField("request", utils.StructToString(data)).WithError(err).Errorf("fail to checking is exist socialmedia")
		return nil, "", err
	}

	if socialmedia == nil {
		uu.Log.WithField("request", utils.StructToString(data)).Errorf("socialmedia is not exist")
		return nil, "Socialmedia Tidak Terdaftar", errors.New("socialmedia not found")
	}

	resultSocialmedia, err := uu.Repo.UpdateSocialmedia(nil, data, ID, userID)
	if err != nil {
		uu.Log.WithField("request", utils.StructToString(data)).WithError(err).Errorf("fail to update socialmedia")
		return nil, "fail to update socialmedia", err
	}

	if resultSocialmedia == nil {
		uu.Log.WithField("request", utils.StructToString(data)).WithError(err).Errorf("fail to update socialmedia")
		return nil, "fail to update socialmedia", err
	}

	return resultSocialmedia, "success update socialmedia", nil
}

func (uu SocialmediaDataUsecase) DeleteSocialmedia(ID int64, userID int64) (*dc.DeleteSocialmediaResponse, string, error) {

	user, err := uu.Repo.GetByIDUserID(ID, userID)
	if err != nil {
		uu.Log.WithField("request", utils.StructToString(ID)).WithError(err).Errorf("fail to checking is exist socialmedia")
		return nil, "", err
	}

	if user == nil {
		uu.Log.WithField("request", utils.StructToString(ID)).Errorf("socialmedia is not exist")
		return nil, "socialmedia Tidak Terdaftar", errors.New("socialmedia not found")
	}

	err = uu.Repo.DeleteSocialmedia(nil, ID, userID)
	if err != nil {
		uu.Log.WithField("request", utils.StructToString(ID)).WithError(err).Errorf("fail to delete socialmedia")
		return nil, "fail to delete socialmedia", err
	}

	result := dc.DeleteSocialmediaResponse{
		Message: "Your socialmedia has been successfully deleted",
	}

	return &result, "success delete socialmedia", nil
}
