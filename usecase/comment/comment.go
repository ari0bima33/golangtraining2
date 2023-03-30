package comment

import (
	"errors"

	dc "final-project/domain/comment"
	"final-project/domain/general"
	"final-project/infra"
	"final-project/repo"
	rc "final-project/repo/comment"
	rp "final-project/repo/photo"
	"final-project/utils"

	"github.com/sirupsen/logrus"
)

type CommentDataUsecaseItf interface {
	GetList() ([]dc.CommentList, string, error)
	CreateComment(data dc.CreateCommentRequest, userID int64) (*dc.CreateCommentResponse, string, error)
	UpdateComment(data dc.UpdateCommentRequest, ID int64, userID int64) (*dc.UpdateCommentResponse, string, error)
	DeleteComment(ID int64, userID int64) (*dc.DeleteCommentResponse, string, error)
}

type CommentDataUsecase struct {
	Repo      rc.CommentDataRepoItf
	RepoPhoto rp.PhotoDataRepoItf
	DBList    *infra.DatabaseList
	Conf      *general.SectionService
	Log       *logrus.Logger
}

func newCommentDataUsecase(r repo.Repo, conf *general.SectionService, logger *logrus.Logger, dbList *infra.DatabaseList) CommentDataUsecase {
	return CommentDataUsecase{
		Repo:      r.Comment.Comment,
		RepoPhoto: r.Photo.Photo,
		Conf:      conf,
		Log:       logger,
		DBList:    dbList,
	}
}

func (uu CommentDataUsecase) GetList() ([]dc.CommentList, string, error) {
	var result []dc.CommentList

	resultCommentList, err := uu.Repo.GetList()
	if err != nil {
		uu.Log.WithField("request", utils.StructToString(nil)).WithError(err).Errorf("fail to get list comment")
		return nil, "fail to get list comment", err
	}

	if len(resultCommentList) > 0 {
		for _, v := range resultCommentList {
			user := dc.User{
				ID:       v.UserID,
				Email:    v.Email,
				Username: v.Username,
			}
			photo := dc.Photo{
				ID:       v.PhotoID,
				Title:    v.PhotoTitle,
				Caption:  v.PhotoCaption,
				PhotoURL: v.PhotoURL,
				UserID:   v.PhotoUserID,
			}
			comment := dc.CommentList{
				ID:        v.ID,
				Message:   v.Message,
				PhotoID:   v.PhotoID,
				UserID:    v.UserID,
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
				User:      user,
				Photo:     photo,
			}
			result = append(result, comment)
		}
	}

	return result, "success get list comment", nil
}

func (uu CommentDataUsecase) CreateComment(data dc.CreateCommentRequest, userID int64) (*dc.CreateCommentResponse, string, error) {

	photo, err := uu.RepoPhoto.GetByID(data.PhotoID)
	if err != nil {
		uu.Log.WithField("request", utils.StructToString(data)).WithError(err).Errorf("fail to checking is exist photo")
		return nil, "", err
	}

	if photo == nil {
		uu.Log.WithField("request", utils.StructToString(data)).Errorf("photo is not exist")
		return nil, "Photo Tidak Terdaftar", errors.New("photo not found")
	}

	resultComment, err := uu.Repo.InsertComment(nil, data, userID)
	if err != nil {
		uu.Log.WithField("request", utils.StructToString(data)).WithError(err).Errorf("fail to Insert Comment")
		return nil, "fail to create comment", err
	}

	if resultComment == nil {
		uu.Log.WithField("request", utils.StructToString(data)).WithError(err).Errorf("fail to Insert Comment")
		return nil, "fail to create comment", err
	}

	return resultComment, "success create comment", nil
}

func (uu CommentDataUsecase) UpdateComment(data dc.UpdateCommentRequest, ID int64, userID int64) (*dc.UpdateCommentResponse, string, error) {

	comment, err := uu.Repo.GetByIDUserID(ID, userID)
	if err != nil {
		uu.Log.WithField("request", utils.StructToString(data)).WithError(err).Errorf("fail to checking is exist comment")
		return nil, "", err
	}

	if comment == nil {
		uu.Log.WithField("request", utils.StructToString(data)).Errorf("comment is not exist")
		return nil, "Comment Tidak Terdaftar", errors.New("comment not found")
	}

	resultComment, err := uu.Repo.UpdateComment(nil, data, ID, userID)
	if err != nil {
		uu.Log.WithField("request", utils.StructToString(data)).WithError(err).Errorf("fail to update comment")
		return nil, "fail to update comment", err
	}

	if resultComment == nil {
		uu.Log.WithField("request", utils.StructToString(data)).WithError(err).Errorf("fail to update comment")
		return nil, "fail to update comment", err
	}

	return resultComment, "success update comment", nil
}

func (uu CommentDataUsecase) DeleteComment(ID int64, userID int64) (*dc.DeleteCommentResponse, string, error) {

	user, err := uu.Repo.GetByIDUserID(ID, userID)
	if err != nil {
		uu.Log.WithField("request", utils.StructToString(ID)).WithError(err).Errorf("fail to checking is exist comment")
		return nil, "", err
	}

	if user == nil {
		uu.Log.WithField("request", utils.StructToString(ID)).Errorf("comment is not exist")
		return nil, "comment Tidak Terdaftar", errors.New("comment not found")
	}

	err = uu.Repo.DeleteComment(nil, ID, userID)
	if err != nil {
		uu.Log.WithField("request", utils.StructToString(ID)).WithError(err).Errorf("fail to delete comment")
		return nil, "fail to delete comment", err
	}

	result := dc.DeleteCommentResponse{
		Message: "Your comment has been successfully deleted",
	}

	return &result, "success delete comment", nil
}
