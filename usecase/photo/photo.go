package photo

import (
	"errors"

	"final-project/domain/general"
	dp "final-project/domain/photo"
	"final-project/infra"
	"final-project/repo"
	rd "final-project/repo/photo"
	"final-project/utils"

	"github.com/sirupsen/logrus"
)

type PhotoDataUsecaseItf interface {
	GetList() ([]dp.PhotoList, string, error)
	CreatePhoto(data dp.CreatePhotoRequest, userID int64) (*dp.CreatePhotoResponse, string, error)
	UpdatePhoto(data dp.UpdatePhotoRequest, ID int64, userID int64) (*dp.UpdatePhotoResponse, string, error)
	DeletePhoto(ID int64, userID int64) (*dp.DeleteUserResponse, string, error)
}

type PhotoDataUsecase struct {
	Repo   rd.PhotoDataRepoItf
	DBList *infra.DatabaseList
	Conf   *general.SectionService
	Log    *logrus.Logger
}

func newPhotoDataUsecase(r repo.Repo, conf *general.SectionService, logger *logrus.Logger, dbList *infra.DatabaseList) PhotoDataUsecase {
	return PhotoDataUsecase{
		Repo:   r.Photo.Photo,
		Conf:   conf,
		Log:    logger,
		DBList: dbList,
	}
}

func (uu PhotoDataUsecase) GetList() ([]dp.PhotoList, string, error) {
	var result []dp.PhotoList

	resultPhotoList, err := uu.Repo.GetList()
	if err != nil {
		uu.Log.WithField("request", utils.StructToString(nil)).WithError(err).Errorf("fail to get list Photo")
		return nil, "fail to get list photo", err
	}

	if len(resultPhotoList) > 0 {
		for _, v := range resultPhotoList {
			user := dp.User{
				Email:    v.Email,
				Username: v.Username,
			}
			photo := dp.PhotoList{
				ID:        v.ID,
				Title:     v.Title,
				Caption:   v.Caption,
				PhotoURL:  v.PhotoURL,
				UserID:    v.UserID,
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
				User:      user,
			}
			result = append(result, photo)
		}
	}

	return result, "success get list photo", nil
}

func (uu PhotoDataUsecase) CreatePhoto(data dp.CreatePhotoRequest, userID int64) (*dp.CreatePhotoResponse, string, error) {

	resultPhoto, err := uu.Repo.InsertPhoto(nil, data, userID)
	if err != nil {
		uu.Log.WithField("request", utils.StructToString(data)).WithError(err).Errorf("fail to Insert Photo")
		return nil, "fail to create photo", err
	}

	if resultPhoto == nil {
		uu.Log.WithField("request", utils.StructToString(data)).WithError(err).Errorf("fail to Insert Photo")
		return nil, "fail to create photo", err
	}

	return resultPhoto, "success create photo", nil
}

func (uu PhotoDataUsecase) UpdatePhoto(data dp.UpdatePhotoRequest, ID int64, userID int64) (*dp.UpdatePhotoResponse, string, error) {

	photo, err := uu.Repo.GetByIDAndUserID(ID, userID)
	if err != nil {
		uu.Log.WithField("request", utils.StructToString(data)).WithError(err).Errorf("fail to checking is exist photo")
		return nil, "", err
	}

	if photo == nil {
		uu.Log.WithField("request", utils.StructToString(data)).Errorf("photo is not exist")
		return nil, "Photo Tidak Terdaftar", errors.New("photo not found")
	}

	if photo.UserID != userID {
		uu.Log.WithField("request", utils.StructToString(data)).Errorf("photo is not exist")
		return nil, "Photo Tidak Terdaftar", errors.New("photo not found")
	}

	resultPhoto, err := uu.Repo.UpdatePhoto(nil, data, ID, userID)
	if err != nil {
		uu.Log.WithField("request", utils.StructToString(data)).WithError(err).Errorf("fail to update photo")
		return nil, "fail to update photo", err
	}

	if resultPhoto == nil {
		uu.Log.WithField("request", utils.StructToString(data)).WithError(err).Errorf("fail to update photo")
		return nil, "fail to update photo", err
	}

	return resultPhoto, "success update photo", nil
}

func (uu PhotoDataUsecase) DeletePhoto(ID int64, userID int64) (*dp.DeleteUserResponse, string, error) {

	photo, err := uu.Repo.GetByIDAndUserID(ID, userID)
	if err != nil {
		uu.Log.WithField("request", utils.StructToString(ID)).WithError(err).Errorf("fail to checking is exist photo")
		return nil, "", err
	}

	if photo == nil {
		uu.Log.WithField("request", utils.StructToString(ID)).Errorf("photo is not exist")
		return nil, "photo Tidak Terdaftar", errors.New("photo not found")
	}

	if photo.UserID != userID {
		uu.Log.WithField("request", utils.StructToString(ID)).Errorf("photo is not exist")
		return nil, "Photo Tidak Terdaftar", errors.New("photo not found")
	}

	err = uu.Repo.DeletePhoto(nil, ID, userID)
	if err != nil {
		uu.Log.WithField("request", utils.StructToString(ID)).WithError(err).Errorf("fail to delete photo")
		return nil, "fail to delete photo", err
	}

	result := dp.DeleteUserResponse{
		Message: "Your photo has been successfully deleted",
	}

	return &result, "success delete photo", nil
}
