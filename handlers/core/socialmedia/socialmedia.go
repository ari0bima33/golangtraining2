package socialmedia

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	cg "final-project/constants/general"
	dg "final-project/domain/general"
	dc "final-project/domain/socialmedia"
	"final-project/handlers"
	"final-project/usecase"
	uc "final-project/usecase/socialmedia"

	"github.com/sirupsen/logrus"
	"gopkg.in/dealancer/validate.v2"
)

type SocialmediaDataHandler struct {
	Usecase uc.SocialmediaDataUsecaseItf
	conf    *dg.SectionService
	log     *logrus.Logger
}

func newSocialmediaHandler(uc usecase.Usecase, conf *dg.SectionService, logger *logrus.Logger) SocialmediaDataHandler {
	return SocialmediaDataHandler{
		Usecase: uc.Socialmedia.Socialmedia,
		conf:    conf,
		log:     logger,
	}
}

func (ch SocialmediaDataHandler) CreateSocialmedia(res http.ResponseWriter, req *http.Request) {
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

	var param dc.CreateSocialmediaRequest

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

	resultSocialmedia, message, err := ch.Usecase.CreateSocialmedia(param, values.UserID)
	if err != nil {
		if message == "" {
			message = "fail to create socialmedia"
		}

		respData.Message = message
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	respData = &handlers.ResponseData{
		Status:  cg.Success,
		Message: message,
	}

	handlers.WriteResponse(res, resultSocialmedia, http.StatusOK)

}

func (ch SocialmediaDataHandler) GetList(res http.ResponseWriter, req *http.Request) {
	respData := &handlers.ResponseData{
		Status: cg.Fail,
	}

	resultList, message, err := ch.Usecase.GetList()
	if err != nil {
		if message == "" {
			message = "fail to get list socialmedia"
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

func (ch SocialmediaDataHandler) UpdateSocialmedia(res http.ResponseWriter, req *http.Request) {
	respData := &handlers.ResponseData{
		Status: cg.Fail,
	}

	paramid := req.FormValue("socialMediaId")
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

	var param dc.UpdateSocialmediaRequest

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

	resultUpdate, message, err := ch.Usecase.UpdateSocialmedia(param, id, values.UserID)
	if err != nil {
		if message == "" {
			message = "fail to update socialmedia"
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

func (ch SocialmediaDataHandler) DeleteSocialmedia(res http.ResponseWriter, req *http.Request) {
	respData := &handlers.ResponseData{
		Status: cg.Fail,
	}

	paramid := req.FormValue("socialMediaId")
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

	resultDelete, message, err := ch.Usecase.DeleteSocialmedia(id, values.UserID)
	if err != nil {
		if message == "" {
			message = "fail to delete socialmedia"
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
