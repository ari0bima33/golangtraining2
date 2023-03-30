package socialmedia

import (
	"database/sql"
	"fmt"
	"time"

	dc "final-project/domain/socialmedia"
	"final-project/infra"
)

type SocialmediaDataRepo struct {
	DBList *infra.DatabaseList
}

func newSocialmediaDataRepo(dbList *infra.DatabaseList) SocialmediaDataRepo {
	return SocialmediaDataRepo{
		DBList: dbList,
	}
}

const (
	uqSelectSocialmedia = `
	SELECT 
	a.id id, 
	a.name "name",
	a.social_media_url social_media_url,
	a.user_id user_id,  
	a.created_at created_at, 
	a.updated_at updated_at,
	COALESCE(u.email,'') email,
	COALESCE(u.username,'') username,
	COALESCE(p.photo_url,'') profile_image_url
	FROM public.socialmedia a
	LEFT JOIN public.user u on a.user_id = u.id
	LEFT JOIN LATERAL (
		SELECT photo_url FROM public.photo
		WHERE user_id = a.user_id
		LIMIT 1
	)as p on true`

	uqInserSocialmedia = `
	INSERT INTO public.socialmedia (
		name,
		social_media_url,
		user_id,
		created_at
	) VALUES (
		?, ?, ?, ?
	)
	RETURNING id,name,social_media_url,user_id,created_at`

	uqUpdateSocialmedia = `
	UPDATE 
		public.socialmedia a
	SET
		name = ?,
		social_media_url = ?,
		updated_at = NOW()`

	uqDeleteSocialmedia = `
		DELETE FROM 
			public.socialmedia a`

	uqWhere = `
	WHERE`

	uqAnd = `
	AND`

	uqFilterID = `
		a.id = ?`

	uqFilterUserID = `
		a.user_id = ?`

	uqReturningSocialmedia = `
	RETURNING id,name,social_media_url,user_id,created_at`
)

type SocialmediaDataRepoItf interface {
	GetList() ([]dc.Socialmedia, error)
	GetByID(ID int64) (*dc.Socialmedia, error)
	GetByIDUserID(ID int64, userID int64) (*dc.Socialmedia, error)
	InsertSocialmedia(tx *sql.Tx, data dc.CreateSocialmediaRequest, userID int64) (*dc.CreateSocialmediaResponse, error)
	UpdateSocialmedia(tx *sql.Tx, data dc.UpdateSocialmediaRequest, ID int64, UserID int64) (*dc.UpdateSocialmediaResponse, error)
	DeleteSocialmedia(tx *sql.Tx, ID int64, UserID int64) error
}

func (ur SocialmediaDataRepo) GetList() ([]dc.Socialmedia, error) {
	var result []dc.Socialmedia

	q := uqSelectSocialmedia
	query, args, err := ur.DBList.Backend.Read.In(q)
	if err != nil {
		return nil, err
	}

	query = ur.DBList.Backend.Read.Rebind(query)
	err = ur.DBList.Backend.Read.Select(&result, query, args...)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (ur SocialmediaDataRepo) GetByID(ID int64) (*dc.Socialmedia, error) {
	var res dc.Socialmedia

	q := fmt.Sprintf("%s %s %s", uqSelectSocialmedia, uqWhere, uqFilterID)
	query, args, err := ur.DBList.Backend.Read.In(q, ID)
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

func (ur SocialmediaDataRepo) GetByIDUserID(ID int64, userID int64) (*dc.Socialmedia, error) {
	var res dc.Socialmedia

	q := fmt.Sprintf("%s %s %s %s %s", uqSelectSocialmedia, uqWhere, uqFilterID, uqAnd, uqFilterUserID)
	query, args, err := ur.DBList.Backend.Read.In(q, ID, userID)
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

func (ur SocialmediaDataRepo) InsertSocialmedia(tx *sql.Tx, data dc.CreateSocialmediaRequest, userID int64) (*dc.CreateSocialmediaResponse, error) {
	param := make([]interface{}, 0)

	param = append(param, data.Name)
	param = append(param, data.SocialMediaUrl)
	param = append(param, userID)
	param = append(param, time.Now().UTC())

	query, args, err := ur.DBList.Backend.Write.In(uqInserSocialmedia, param...)
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

	var socialmedia = dc.CreateSocialmediaResponse{}
	err = res.Scan(&socialmedia.ID, &socialmedia.Name, &socialmedia.SocialMediaUrl, &socialmedia.UserID, &socialmedia.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &socialmedia, nil
}

func (ur SocialmediaDataRepo) UpdateSocialmedia(tx *sql.Tx, data dc.UpdateSocialmediaRequest, ID int64, UserID int64) (*dc.UpdateSocialmediaResponse, error) {
	var err error

	q := fmt.Sprintf("%s %s %s %s %s %s", uqUpdateSocialmedia, uqWhere, uqFilterID, uqAnd, uqFilterUserID, uqReturningSocialmedia)

	query, args, err := ur.DBList.Backend.Read.In(q, data.Name, data.SocialMediaUrl, ID, UserID)
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

	var socialmedia = dc.UpdateSocialmediaResponse{}
	err = res.Scan(&socialmedia.ID, &socialmedia.Name, &socialmedia.SocialMediaUrl, &socialmedia.UserID, &socialmedia.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &socialmedia, nil
}

func (ur SocialmediaDataRepo) DeleteSocialmedia(tx *sql.Tx, ID int64, UserID int64) error {
	var err error

	q := fmt.Sprintf("%s %s %s %s %s", uqDeleteSocialmedia, uqWhere, uqFilterID, uqAnd, uqFilterUserID)

	query, args, err := ur.DBList.Backend.Read.In(q, ID, UserID)
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
