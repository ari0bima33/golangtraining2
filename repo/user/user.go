package user

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	du "final-project/domain/user"
	"final-project/infra"
)

type UserDataRepo struct {
	DBList *infra.DatabaseList
}

func newUserDataRepo(dbList *infra.DatabaseList) UserDataRepo {
	return UserDataRepo{
		DBList: dbList,
	}
}

const (
	uqSelectUser = `
	SELECT
		id,
		username,
		email,
		password,
		age,
		created_at,
		updated_at
	FROM
		public.user`

	uqInsertUser = `
	INSERT INTO public.user (
		username,
		email,
		password,
		age,
		created_at
	) VALUES (
		?, ?, ?, ?, ?
	)
	RETURNING id,username,email,password,age,created_at`

	uqUpdateUser = `
	UPDATE 
		public.user
	SET
		username = ?,
		email = ?,
		updated_at = NOW()`

	uqDeleteUser = `
		DELETE FROM 
			public.user`

	uqWhere = `
	WHERE`

	uqFilterUserID = `
		id = ?`

	uqFilterLikeUserName = `
		lower(username) LIKE ?`

	uqSelectExist = `
		SELECT EXISTS`

	uqFilterEmailFilter = `
		email = ?`
	uqFilterUsernameFilter = `
		username = ?`
	uqNotUserIDFilter = ` 
	AND id <> ?`

	uqReturningUser = `
	RETURNING id,username,email,password,age,created_at,updated_at`
)

type UserDataRepoItf interface {
	GetByID(userID int64) (*du.User, error)
	GetByUsername(userName string) (*du.User, error)
	GetByEmail(email string) (*du.User, error)
	UpdateUser(tx *sql.Tx, username string, email string, userID int64) (*du.User, error)
	InsertUser(tx *sql.Tx, data du.CreateUser) (*du.User, error)
	IsExistUserEmail(email string) (bool, error)
	IsExistUsername(username string) (bool, error)
	IsExistUserEmailAndUserID(email string, userID int64) (bool, error)
	IsExistUsernameAndUserID(username string, userID int64) (bool, error)
	DeleteUser(tx *sql.Tx, userID int64) error
}

func (ur UserDataRepo) GetByID(userID int64) (*du.User, error) {
	var res du.User

	q := fmt.Sprintf("%s%s%s", uqSelectUser, uqWhere, uqFilterUserID)
	query, args, err := ur.DBList.Backend.Read.In(q, userID)
	if err != nil {
		return nil, err
	}

	query = ur.DBList.Backend.Read.Rebind(query)
	err = ur.DBList.Backend.Read.Get(&res, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if res.ID == 0 {
		return nil, nil
	}

	return &res, nil
}

func (ur UserDataRepo) GetByUsername(userName string) (*du.User, error) {
	var res *du.User

	q := fmt.Sprintf("%s%s%s", uqSelectUser, uqWhere, uqFilterLikeUserName)
	query, args, err := ur.DBList.Backend.Read.In(q, strings.ToLower(userName))
	if err != nil {
		return nil, err
	}

	query = ur.DBList.Backend.Read.Rebind(query)
	err = ur.DBList.Backend.Read.Get(&res, query, args...)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (ur UserDataRepo) GetByEmail(email string) (*du.User, error) {
	var res du.User

	q := fmt.Sprintf("%s%s%s", uqSelectUser, uqWhere, uqFilterEmailFilter)
	query, args, err := ur.DBList.Backend.Read.In(q, email)
	if err != nil {
		return nil, err
	}

	query = ur.DBList.Backend.Read.Rebind(query)
	err = ur.DBList.Backend.Read.Get(&res, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if res.ID == 0 {
		return nil, nil
	}

	return &res, nil
}

func (ur UserDataRepo) InsertUser(tx *sql.Tx, data du.CreateUser) (*du.User, error) {
	param := make([]interface{}, 0)

	param = append(param, strings.Title(strings.ToLower(data.Username)))
	param = append(param, data.Email)
	param = append(param, data.Password)
	param = append(param, data.Age)
	param = append(param, time.Now().UTC())

	query, args, err := ur.DBList.Backend.Write.In(uqInsertUser, param...)
	if err != nil {
		return nil, err
	}

	query = ur.DBList.Backend.Write.Rebind(query)

	var res *sql.Row
	if tx == nil {
		res = ur.DBList.Backend.Write.QueryRow(query, args...)
	} else {
		res = tx.QueryRow(query, args...)
	}

	if err != nil {
		return nil, err
	}

	err = res.Err()
	if err != nil {
		return nil, err
	}

	var user = du.User{}
	err = res.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Age, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur UserDataRepo) UpdateUser(tx *sql.Tx, username string, email string, userID int64) (*du.User, error) {
	var err error

	q := fmt.Sprintf("%s %s%s %s", uqUpdateUser, uqWhere, uqFilterUserID, uqReturningUser)

	query, args, err := ur.DBList.Backend.Read.In(q, username, email, userID)
	if err != nil {
		return nil, err
	}

	query = ur.DBList.Backend.Write.Rebind(query)

	var res *sql.Row
	if tx == nil {
		res = ur.DBList.Backend.Write.QueryRow(query, args...)
	} else {
		res = tx.QueryRow(query, args...)
	}

	if err != nil {
		return nil, err
	}

	err = res.Err()
	if err != nil {
		return nil, err
	}

	var user = du.User{}
	err = res.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Age, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur UserDataRepo) IsExistUserEmail(email string) (bool, error) {
	var isExist bool

	q := fmt.Sprintf("%s(%s%s%s)", uqSelectExist, uqSelectUser, uqWhere, uqFilterEmailFilter)

	query, args, err := ur.DBList.Backend.Read.In(q, email)
	if err != nil {
		return isExist, err
	}

	query = ur.DBList.Backend.Read.Rebind(query)
	err = ur.DBList.Backend.Read.Get(&isExist, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return isExist, err
	}

	return isExist, nil
}

func (ur UserDataRepo) IsExistUsername(username string) (bool, error) {
	var isExist bool

	q := fmt.Sprintf("%s(%s%s%s)", uqSelectExist, uqSelectUser, uqWhere, uqFilterUsernameFilter)

	query, args, err := ur.DBList.Backend.Read.In(q, username)
	if err != nil {
		return isExist, err
	}

	query = ur.DBList.Backend.Read.Rebind(query)
	err = ur.DBList.Backend.Read.Get(&isExist, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return isExist, err
	}

	return isExist, nil
}

func (ur UserDataRepo) IsExistUserEmailAndUserID(email string, userID int64) (bool, error) {
	var isExist bool

	q := fmt.Sprintf("%s(%s%s%s%s)", uqSelectExist, uqSelectUser, uqWhere, uqFilterEmailFilter, uqNotUserIDFilter)

	query, args, err := ur.DBList.Backend.Read.In(q, email, userID)
	if err != nil {
		return isExist, err
	}

	query = ur.DBList.Backend.Read.Rebind(query)
	err = ur.DBList.Backend.Read.Get(&isExist, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return isExist, err
	}

	return isExist, nil
}

func (ur UserDataRepo) IsExistUsernameAndUserID(username string, userID int64) (bool, error) {
	var isExist bool

	q := fmt.Sprintf("%s(%s%s%s%s)", uqSelectExist, uqSelectUser, uqWhere, uqFilterUsernameFilter, uqNotUserIDFilter)

	query, args, err := ur.DBList.Backend.Read.In(q, username, userID)
	if err != nil {
		return isExist, err
	}

	query = ur.DBList.Backend.Read.Rebind(query)
	err = ur.DBList.Backend.Read.Get(&isExist, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return isExist, err
	}

	return isExist, nil
}

func (ur UserDataRepo) DeleteUser(tx *sql.Tx, userID int64) error {
	var err error

	q := fmt.Sprintf("%s%s%s", uqDeleteUser, uqWhere, uqFilterUserID)

	query, args, err := ur.DBList.Backend.Read.In(q, userID)
	if err != nil {
		return err
	}

	query = ur.DBList.Backend.Write.Rebind(query)
	_, err = ur.DBList.Backend.Write.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}
