package user

import (
	"errors"
	"fmt"

	"final-project/domain/general"
	du "final-project/domain/user"
	"final-project/infra"
	"final-project/repo"
	ru "final-project/repo/user"
	"final-project/utils"

	"github.com/sirupsen/logrus"
)

type UserDataUsecaseItf interface {
	RegisterUser(data du.CreateUserRequest) (*du.CreateUserResponse, string, error)
	UserLogin(data du.UserLoginRequest) (*du.UserLoginResponse, string, error)
	UpdateUser(userID int64, data du.UpdateUserRequest) (*du.UpdateUserResponse, string, error)
	DeleteUser(userID int64) (*du.DeleteUserResponse, string, error)
}

type UserDataUsecase struct {
	Repo   ru.UserDataRepoItf
	DBList *infra.DatabaseList
	Conf   *general.SectionService
	Log    *logrus.Logger
}

func newUserDataUsecase(r repo.Repo, conf *general.SectionService, logger *logrus.Logger, dbList *infra.DatabaseList) UserDataUsecase {
	return UserDataUsecase{
		Repo:   r.User.User,
		Conf:   conf,
		Log:    logger,
		DBList: dbList,
	}
}

func (uu UserDataUsecase) RegisterUser(data du.CreateUserRequest) (*du.CreateUserResponse, string, error) {
	isEmailValid := utils.EmailValidator(data.Email)
	if !isEmailValid {
		uu.Log.WithField("request", utils.StructToString(data)).Errorf("Email not valid")
		return nil, "Email tidak valid", errors.New("Email not valid")
	}

	isExistEmail, err := uu.Repo.IsExistUserEmail(data.Email)
	if err != nil {
		uu.Log.WithField("request", utils.StructToString(data)).WithError(err).Errorf("fail to checking is exist user")
		return nil, "", err
	}

	if isExistEmail {
		uu.Log.WithField("request", utils.StructToString(data)).Errorf("user is not exist")
		return nil, "Email Sudah Terdaftar", errors.New("user already exist")
	}

	IsExistUsername, err := uu.Repo.IsExistUsername(data.Username)
	if err != nil {
		uu.Log.WithField("request", utils.StructToString(data)).WithError(err).Errorf("fail to checking is exist user")
		return nil, "", err
	}

	if IsExistUsername {
		uu.Log.WithField("request", utils.StructToString(data)).Errorf("username already exist")
		return nil, "Username Sudah Terdaftar", errors.New("username already exist")
	}

	generatePassword, err := utils.GeneratePassword(data.Password)
	if err != nil {
		uu.Log.WithField("request", utils.StructToString(data)).WithError(err).Errorf("fail to Generate Password")
		return nil, "", err
	}

	insertuser := du.CreateUser{Email: data.Email, Age: data.Age, Username: data.Username, Password: generatePassword}

	var resultuser *du.User
	resultuser, err = uu.Repo.InsertUser(nil, insertuser)
	if err != nil {
		uu.Log.WithField("request", utils.StructToString(data)).WithError(err).Errorf("fail to Insert User")
		return nil, "fail to register user", err
	}

	if resultuser == nil {
		uu.Log.WithField("request", utils.StructToString(data)).WithError(err).Errorf("fail to Insert User")
		return nil, "fail to register user", err
	}

	result := du.CreateUserResponse{Age: resultuser.Age, Email: resultuser.Email, Username: resultuser.Username, ID: resultuser.ID}

	return &result, "success register user", nil
}

func (uu UserDataUsecase) UserLogin(data du.UserLoginRequest) (*du.UserLoginResponse, string, error) {

	isEmailValid := utils.EmailValidator(data.Email)
	if !isEmailValid {
		uu.Log.WithField("request", utils.StructToString(data)).Errorf("Email not valid")
		return nil, "Email tidak valid", errors.New("Email not valid")
	}

	resultUser, err := uu.Repo.GetByEmail(data.Email)
	if err != nil {
		uu.Log.WithField("request", utils.StructToString(data)).WithError(err).Errorf("fail to checking is exist user")
		return nil, "", err
	}

	if resultUser == nil {
		uu.Log.WithField("request", utils.StructToString(data)).Errorf("user is not exist")
		return nil, "User Tidak Terdaftar", errors.New("user not found")
	}

	comparePassword, err := utils.ComparePassword(resultUser.Password, data.Password)
	if err != nil {
		uu.Log.WithField("request", utils.StructToString(data)).WithError(err).Errorf("fail to Compare Password")
		return nil, "", err
	}

	if !comparePassword {
		uu.Log.WithField("request", utils.StructToString(data)).Errorf("password is wrong")
		return nil, "Password Salah", errors.New("password is wrong")
	}

	session, err := utils.GetEncrypt([]byte(uu.Conf.App.SecretKey), fmt.Sprintf("%v", 1))
	if err != nil {
		uu.Log.WithField("user id", resultUser.ID).WithError(err).Error("fail to get session token data from infra")
		return nil, "", err
	}

	accessToken, _, err := utils.GenerateJWTUserID(session, resultUser.ID)
	if err != nil {
		uu.Log.WithField("user id", resultUser.ID).WithError(err).Error("fail to get token data from infra")
		return nil, "", err
	}

	result := du.UserLoginResponse{Token: accessToken}

	return &result, "success login user", nil
}

func (uu UserDataUsecase) UpdateUser(userID int64, data du.UpdateUserRequest) (*du.UpdateUserResponse, string, error) {

	user, err := uu.Repo.GetByID(userID)
	if err != nil {
		uu.Log.WithField("request", utils.StructToString(data)).WithError(err).Errorf("fail to checking is exist user")
		return nil, "", err
	}

	if user == nil {
		uu.Log.WithField("request", utils.StructToString(data)).Errorf("user is not exist")
		return nil, "User Tidak Terdaftar", errors.New("user not found")
	}

	isEmailValid := utils.EmailValidator(data.Email)
	if !isEmailValid {
		uu.Log.WithField("request", utils.StructToString(data)).Errorf("Email not valid")
		return nil, "Email tidak valid", errors.New("Email not valid")
	}

	IsExistEmail, err := uu.Repo.IsExistUserEmailAndUserID(data.Email, userID)
	if err != nil {
		uu.Log.WithField("request", utils.StructToString(data)).WithError(err).Errorf("fail to checking is exist user")
		return nil, "", err
	}

	if IsExistEmail {
		uu.Log.WithField("request", utils.StructToString(data)).Errorf("email already exist")
		return nil, "Email Sudah Terdaftar", errors.New("email already exist")
	}

	IsExistUsername, err := uu.Repo.IsExistUsernameAndUserID(data.Username, userID)
	if err != nil {
		uu.Log.WithField("request", utils.StructToString(data)).WithError(err).Errorf("fail to checking is exist user")
		return nil, "", err
	}

	if IsExistUsername {
		uu.Log.WithField("request", utils.StructToString(data)).Errorf("username already exist")
		return nil, "Username Sudah Terdaftar", errors.New("username already exist")
	}
	var resultUser *du.User
	resultUser, err = uu.Repo.UpdateUser(nil, data.Username, data.Email, userID)
	if err != nil {
		uu.Log.WithField("request", utils.StructToString(data)).WithError(err).Errorf("fail to update user")
		return nil, "fail to update user", err
	}

	if resultUser == nil {
		uu.Log.WithField("request", utils.StructToString(data)).WithError(err).Errorf("fail to update user")
		return nil, "fail to update user", err
	}

	result := du.UpdateUserResponse{
		Age:       resultUser.Age,
		Email:     resultUser.Email,
		Username:  resultUser.Username,
		ID:        resultUser.ID,
		UpdatedAt: resultUser.UpdatedAt,
	}

	return &result, "success update user", nil
}

func (uu UserDataUsecase) DeleteUser(userID int64) (*du.DeleteUserResponse, string, error) {

	user, err := uu.Repo.GetByID(userID)
	if err != nil {
		uu.Log.WithField("request", utils.StructToString(userID)).WithError(err).Errorf("fail to checking is exist user")
		return nil, "", err
	}

	if user == nil {
		uu.Log.WithField("request", utils.StructToString(userID)).Errorf("user is not exist")
		return nil, "User Tidak Terdaftar", errors.New("user not found")
	}

	err = uu.Repo.DeleteUser(nil, userID)
	if err != nil {
		uu.Log.WithField("request", utils.StructToString(userID)).WithError(err).Errorf("fail to delete user")
		return nil, "fail to delete user", err
	}

	result := du.DeleteUserResponse{
		Message: "Your account has been successfully deleted",
	}

	return &result, "success delete user", nil
}
