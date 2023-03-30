package photo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	cg "final-project/constants/general"
	dg "final-project/domain/general"
	dp "final-project/domain/photo"
	"final-project/handlers"
	"final-project/usecase"
	up "final-project/usecase/photo"

	"github.com/sirupsen/logrus"
	"gopkg.in/dealancer/validate.v2"
)

type PhotoDataHandler struct {
	Usecase up.PhotoDataUsecaseItf
	conf    *dg.SectionService
	log     *logrus.Logger
}

func newPhotoHandler(uc usecase.Usecase, conf *dg.SectionService, logger *logrus.Logger) PhotoDataHandler {
	return PhotoDataHandler{
		Usecase: uc.Photo.Photo,
		conf:    conf,
		log:     logger,
	}
}

func (ch PhotoDataHandler) CreatePhoto(res http.ResponseWriter, req *http.Request) {
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

	var param dp.CreatePhotoRequest

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

	resultPhoto, message, err := ch.Usecase.CreatePhoto(param, values.UserID)
	if err != nil {
		if message == "" {
			message = "fail to create photo"
		}

		respData.Message = message
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	respData = &handlers.ResponseData{
		Status:  cg.Success,
		Message: message,
	}

	handlers.WriteResponse(res, resultPhoto, http.StatusOK)

}

func (ch PhotoDataHandler) GetList(res http.ResponseWriter, req *http.Request) {
	respData := &handlers.ResponseData{
		Status: cg.Fail,
	}

	resultList, message, err := ch.Usecase.GetList()
	if err != nil {
		if message == "" {
			message = "fail to get list photo"
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

func (ch PhotoDataHandler) UpdatePhoto(res http.ResponseWriter, req *http.Request) {
	respData := &handlers.ResponseData{
		Status: cg.Fail,
	}

	paramid := req.FormValue("photoId")
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

	var param dp.UpdatePhotoRequest

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

	resultUpdate, message, err := ch.Usecase.UpdatePhoto(param, id, values.UserID)
	if err != nil {
		if message == "" {
			message = "fail to update photo"
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

func (ch PhotoDataHandler) DeletePhoto(res http.ResponseWriter, req *http.Request) {
	respData := &handlers.ResponseData{
		Status: cg.Fail,
	}

	paramid := req.FormValue("photoId")
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

	resultDelete, message, err := ch.Usecase.DeletePhoto(id, values.UserID)
	if err != nil {
		if message == "" {
			message = "fail to delete photo"
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
