package comment

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	cg "final-project/constants/general"
	dc "final-project/domain/comment"
	dg "final-project/domain/general"
	"final-project/handlers"
	"final-project/usecase"
	uc "final-project/usecase/comment"

	"github.com/sirupsen/logrus"
	"gopkg.in/dealancer/validate.v2"
)

type CommentDataHandler struct {
	Usecase uc.CommentDataUsecaseItf
	conf    *dg.SectionService
	log     *logrus.Logger
}

func newCommentHandler(uc usecase.Usecase, conf *dg.SectionService, logger *logrus.Logger) CommentDataHandler {
	return CommentDataHandler{
		Usecase: uc.Comment.Comment,
		conf:    conf,
		log:     logger,
	}
}

func (ch CommentDataHandler) CreateComment(res http.ResponseWriter, req *http.Request) {
	respData := &handlers.ResponseData{
		Status: cg.Fail,
	}

	contextValues := req.Context().Value("values")
	values, ok := contextValues.(dg.ContextValue)
	if !ok {
		respData.Message = cg.HandlerErrorRequestDataNotValid
		handlers.WriteResponse(res, respData, http.StatusBadRequest)
		return
	}

	var param dc.CreateCommentRequest

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		respData.Message = cg.HandlerErrorRequestDataEmpty
		handlers.WriteResponse(res, respData, http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(reqBody, &param)
	if err != nil {
		respData.Message = cg.HandlerErrorRequestDataNotValid
		handlers.WriteResponse(res, respData, http.StatusBadRequest)
		return
	}

	err = validate.Validate(param)
	if err != nil {
		respData.Message = cg.HandlerErrorRequestDataFormatInvalid
		handlers.WriteResponse(res, respData, http.StatusBadRequest)
		return
	}

	resultComment, message, err := ch.Usecase.CreateComment(param, values.UserID)
	if err != nil {
		if message == "" {
			message = "fail to create comment"
		}

		respData.Message = message
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	respData = &handlers.ResponseData{
		Status:  cg.Success,
		Message: message,
	}

	handlers.WriteResponse(res, resultComment, http.StatusOK)

}

func (ch CommentDataHandler) GetList(res http.ResponseWriter, req *http.Request) {
	respData := &handlers.ResponseData{
		Status: cg.Fail,
	}

	resultList, message, err := ch.Usecase.GetList()
	if err != nil {
		if message == "" {
			message = "fail to get list comment"
		}

		respData.Message = message
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	respData = &handlers.ResponseData{
		Status:  cg.Success,
		Message: message,
	}

	handlers.WriteResponse(res, resultList, http.StatusOK)

}

func (ch CommentDataHandler) UpdateComment(res http.ResponseWriter, req *http.Request) {
	respData := &handlers.ResponseData{
		Status: cg.Fail,
	}

	paramid := req.FormValue("commentId")
	id, err := strconv.ParseInt(paramid, 10, 64)
	if err != nil {
		respData.Message = cg.HandlerErrorRequestDataEmpty
		handlers.WriteResponse(res, respData, http.StatusBadRequest)
		return
	}

	contextValues := req.Context().Value("values")
	values, ok := contextValues.(dg.ContextValue)
	if !ok {
		respData.Message = cg.HandlerErrorRequestDataNotValid
		handlers.WriteResponse(res, respData, http.StatusBadRequest)
		return
	}

	var param dc.UpdateCommentRequest

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		respData.Message = cg.HandlerErrorRequestDataEmpty
		handlers.WriteResponse(res, respData, http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(reqBody, &param)
	if err != nil {
		respData.Message = cg.HandlerErrorRequestDataNotValid
		handlers.WriteResponse(res, respData, http.StatusBadRequest)
		return
	}

	err = validate.Validate(param)
	if err != nil {
		respData.Message = cg.HandlerErrorRequestDataFormatInvalid
		handlers.WriteResponse(res, respData, http.StatusBadRequest)
		return
	}

	resultUpdate, message, err := ch.Usecase.UpdateComment(param, id, values.UserID)
	if err != nil {
		if message == "" {
			message = "fail to update comment"
		}

		respData.Message = message
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	respData = &handlers.ResponseData{
		Status:  cg.Success,
		Message: message,
	}

	handlers.WriteResponse(res, resultUpdate, http.StatusOK)

}

func (ch CommentDataHandler) DeleteComment(res http.ResponseWriter, req *http.Request) {
	respData := &handlers.ResponseData{
		Status: cg.Fail,
	}

	paramid := req.FormValue("commentId")
	id, err := strconv.ParseInt(paramid, 10, 64)
	if err != nil {
		respData.Message = cg.HandlerErrorRequestDataEmpty
		handlers.WriteResponse(res, respData, http.StatusBadRequest)
		return
	}

	contextValues := req.Context().Value("values")
	values, ok := contextValues.(dg.ContextValue)
	if !ok {
		respData.Message = cg.HandlerErrorRequestDataNotValid
		handlers.WriteResponse(res, respData, http.StatusBadRequest)
		return
	}

	resultDelete, message, err := ch.Usecase.DeleteComment(id, values.UserID)
	if err != nil {
		if message == "" {
			message = "fail to delete comment"
		}

		respData.Message = message
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	respData = &handlers.ResponseData{
		Status:  cg.Success,
		Message: message,
	}

	handlers.WriteResponse(res, resultDelete, http.StatusOK)

}
