package user

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	cg "final-project/constants/general"
	"final-project/domain/general"
	du "final-project/domain/user"
	"final-project/handlers"
	"final-project/usecase"
	uu "final-project/usecase/user"

	"github.com/sirupsen/logrus"
	"gopkg.in/dealancer/validate.v2"
)

type UserDataHandler struct {
	Usecase uu.UserDataUsecaseItf
	conf    *general.SectionService
	log     *logrus.Logger
}

func newUserHandler(uc usecase.Usecase, conf *general.SectionService, logger *logrus.Logger) UserDataHandler {
	return UserDataHandler{
		Usecase: uc.User.User,
		conf:    conf,
		log:     logger,
	}
}

func (ch UserDataHandler) RegisterUser(res http.ResponseWriter, req *http.Request) {
	respData := &handlers.ResponseData{
		Status: cg.Fail,
	}

	var param du.CreateUserRequest

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

	resultUser, message, err := ch.Usecase.RegisterUser(param)
	if err != nil {
		if message == "" {
			message = "fail to register user"
		}

		respData.Message = message
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	respData = &handlers.ResponseData{
		Status:  cg.Success,
		Message: message,
	}

	handlers.WriteResponse(res, resultUser, http.StatusOK)

}

func (ch UserDataHandler) UserLogin(res http.ResponseWriter, req *http.Request) {
	respData := &handlers.ResponseData{
		Status: cg.Fail,
	}

	var param du.UserLoginRequest

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

	resultLogin, message, err := ch.Usecase.UserLogin(param)
	if err != nil {
		if message == "" {
			message = "fail to login user"
		}

		respData.Message = message
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	respData = &handlers.ResponseData{
		Status:  cg.Success,
		Message: message,
	}

	handlers.WriteResponse(res, resultLogin, http.StatusOK)

}

func (ch UserDataHandler) UpdateUser(res http.ResponseWriter, req *http.Request) {
	respData := &handlers.ResponseData{
		Status: cg.Fail,
	}

	paramid := req.FormValue("userid")
	id, err := strconv.ParseInt(paramid, 10, 64)
	if err != nil {
		respData.Message = cg.HandlerErrorRequestDataEmpty
		handlers.WriteResponse(res, respData, http.StatusBadRequest)
		return
	}

	var param du.UpdateUserRequest

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

	resultUpdate, message, err := ch.Usecase.UpdateUser(id, param)
	if err != nil {
		if message == "" {
			message = "fail to update user"
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

func (ch UserDataHandler) DeleteUser(res http.ResponseWriter, req *http.Request) {
	respData := &handlers.ResponseData{
		Status: cg.Fail,
	}

	paramid := req.FormValue("userid")
	id, err := strconv.ParseInt(paramid, 10, 64)
	if err != nil {
		respData.Message = cg.HandlerErrorRequestDataEmpty
		handlers.WriteResponse(res, respData, http.StatusBadRequest)
		return
	}

	resultDelete, message, err := ch.Usecase.DeleteUser(id)
	if err != nil {
		if message == "" {
			message = "fail to update user"
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
