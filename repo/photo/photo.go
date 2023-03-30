package photo

import (
	"database/sql"
	"fmt"
	"time"

	dp "final-project/domain/photo"
	"final-project/infra"
)

type PhotoDataRepo struct {
	DBList *infra.DatabaseList
}

func newPhotoDataRepo(dbList *infra.DatabaseList) PhotoDataRepo {
	return PhotoDataRepo{
		DBList: dbList,
	}
}

const (
	uqSelectPhoto = `
	SELECT 
	p.id id, 
	p.title title, 
	p.caption caption, 
	p.photo_url photo_url, 
	p.user_id user_id,  
	p.created_at created_at, 
	p.updated_at updated_at,
	COALESCE(u.email,'') email,
	COALESCE(u.username,'') username
	FROM public.photo p
	LEFT JOIN public.user u on p.user_id = u.id`

	uqInserPhoto = `
	INSERT INTO public.photo (
		title,
		caption,
		photo_url,
		user_id,
		created_at
	) VALUES (
		?, ?, ?, ?, ?
	)
	RETURNING id,title,caption,photo_url,user_id,created_at`

	uqUpdatePhoto = `
	UPDATE 
		public.photo p
	SET
		title = ?,
		caption = ?,
		photo_url = ?,
		updated_at = NOW()`

	uqDeletePhoto = `
		DELETE FROM 
			public.photo p`

	uqWhere = `
	WHERE`

	uqAnd = `
	And`

	uqFilterID = `
		p.id = ?`

	uqFilterUserID = `
		p.user_id = ?`

	uqReturningPhoto = `
	RETURNING id,title,caption,photo_url,user_id,updated_at`
)

type PhotoDataRepoItf interface {
	GetList() ([]dp.Photo, error)
	GetByID(ID int64) (*dp.Photo, error)
	GetByIDAndUserID(ID int64, userID int64) (*dp.Photo, error)
	InsertPhoto(tx *sql.Tx, data dp.CreatePhotoRequest, userID int64) (*dp.CreatePhotoResponse, error)
	UpdatePhoto(tx *sql.Tx, data dp.UpdatePhotoRequest, ID int64, UserID int64) (*dp.UpdatePhotoResponse, error)
	DeletePhoto(tx *sql.Tx, ID int64, UserID int64) error
}

func (ur PhotoDataRepo) GetList() ([]dp.Photo, error) {
	var result []dp.Photo

	q := uqSelectPhoto
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

func (ur PhotoDataRepo) GetByID(ID int64) (*dp.Photo, error) {
	var res dp.Photo

	q := fmt.Sprintf("%s %s %s", uqSelectPhoto, uqWhere, uqFilterID)
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

func (ur PhotoDataRepo) GetByIDAndUserID(ID int64, userID int64) (*dp.Photo, error) {
	var res dp.Photo

	q := fmt.Sprintf("%s %s %s %s %s", uqSelectPhoto, uqWhere, uqFilterID, uqAnd, uqFilterUserID)
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

func (ur PhotoDataRepo) InsertPhoto(tx *sql.Tx, data dp.CreatePhotoRequest, userID int64) (*dp.CreatePhotoResponse, error) {
	param := make([]interface{}, 0)

	param = append(param, data.Title)
	param = append(param, data.Caption)
	param = append(param, data.PhotoURL)
	param = append(param, userID)
	param = append(param, time.Now().UTC())

	query, args, err := ur.DBList.Backend.Write.In(uqInserPhoto, param...)
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

	var photo = dp.CreatePhotoResponse{}
	err = res.Scan(&photo.ID, &photo.Title, &photo.Caption, &photo.PhotoURL, &photo.UserID, &photo.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &photo, nil
}

func (ur PhotoDataRepo) UpdatePhoto(tx *sql.Tx, data dp.UpdatePhotoRequest, ID int64, UserID int64) (*dp.UpdatePhotoResponse, error) {
	var err error

	q := fmt.Sprintf("%s %s %s %s %s %s", uqUpdatePhoto, uqWhere, uqFilterID, uqAnd, uqFilterUserID, uqReturningPhoto)

	query, args, err := ur.DBList.Backend.Read.In(q, data.Title, data.Caption, data.PhotoURL, ID, UserID)
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

	var photo = dp.UpdatePhotoResponse{}
	err = res.Scan(&photo.ID, &photo.Title, &photo.Caption, &photo.PhotoURL, &photo.UserID, &photo.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &photo, nil
}

func (ur PhotoDataRepo) DeletePhoto(tx *sql.Tx, ID int64, UserID int64) error {
	var err error

	q := fmt.Sprintf("%s %s %s %s %s", uqDeletePhoto, uqWhere, uqFilterID, uqAnd, uqFilterUserID)

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
