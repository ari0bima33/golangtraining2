package comment

import (
	"database/sql"
	"fmt"
	"time"

	dc "final-project/domain/comment"
	"final-project/infra"
)

type CommentDataRepo struct {
	DBList *infra.DatabaseList
}

func newCommentDataRepo(dbList *infra.DatabaseList) CommentDataRepo {
	return CommentDataRepo{
		DBList: dbList,
	}
}

const (
	uqSelectComment = `
	SELECT 
	c.id id, 
	c.photo_id photo_id,
	c.message message,
	c.user_id user_id,  
	c.created_at created_at, 
	c.updated_at updated_at,
	COALESCE(u.email,'') email,
	COALESCE(u.username,'') username,
	COALESCE(p.title,'') photo_title,
	COALESCE(p.caption,'') photo_caption,
	COALESCE(p.photo_url,'') photo_url,
	p.user_id photo_user_id
	FROM public.comment c
	LEFT JOIN public.user u on c.user_id = u.id
	LEFT JOIN public.photo p on c.photo_id = p.id`

	uqInserComment = `
	INSERT INTO public.comment (
		message,
		photo_id,
		user_id,
		created_at
	) VALUES (
		?, ?, ?, ?
	)
	RETURNING id,message,photo_id,user_id,created_at`

	uqUpdateComment = `
	UPDATE 
		public.comment c
	SET
		message = ?,
		updated_at = NOW()`

	uqDeleteComment = `
		DELETE FROM 
			public.comment c`

	uqWhere = `
	WHERE`

	uqAnd = `
	AND`

	uqFilterID = `
		c.id = ?`

	uqFilterUserID = `
		c.user_id = ?`

	uqReturningComment = `
	RETURNING id,message,photo_id,user_id,created_at`
)

type CommentDataRepoItf interface {
	GetList() ([]dc.Comment, error)
	GetByID(ID int64) (*dc.Comment, error)
	GetByIDUserID(ID int64, userID int64) (*dc.Comment, error)
	InsertComment(tx *sql.Tx, data dc.CreateCommentRequest, userID int64) (*dc.CreateCommentResponse, error)
	UpdateComment(tx *sql.Tx, data dc.UpdateCommentRequest, ID int64, UserID int64) (*dc.UpdateCommentResponse, error)
	DeleteComment(tx *sql.Tx, ID int64, UserID int64) error
}

func (ur CommentDataRepo) GetList() ([]dc.Comment, error) {
	var result []dc.Comment

	q := uqSelectComment
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

func (ur CommentDataRepo) GetByID(ID int64) (*dc.Comment, error) {
	var res dc.Comment

	q := fmt.Sprintf("%s %s %s", uqSelectComment, uqWhere, uqFilterID)
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

func (ur CommentDataRepo) GetByIDUserID(ID int64, userID int64) (*dc.Comment, error) {
	var res dc.Comment

	q := fmt.Sprintf("%s %s %s %s %s", uqSelectComment, uqWhere, uqFilterID, uqAnd, uqFilterUserID)
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

func (ur CommentDataRepo) InsertComment(tx *sql.Tx, data dc.CreateCommentRequest, userID int64) (*dc.CreateCommentResponse, error) {
	param := make([]interface{}, 0)

	param = append(param, data.Message)
	param = append(param, data.PhotoID)
	param = append(param, userID)
	param = append(param, time.Now().UTC())

	query, args, err := ur.DBList.Backend.Write.In(uqInserComment, param...)
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

	var comment = dc.CreateCommentResponse{}
	err = res.Scan(&comment.ID, &comment.Message, &comment.PhotoID, &comment.UserID, &comment.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &comment, nil
}

func (ur CommentDataRepo) UpdateComment(tx *sql.Tx, data dc.UpdateCommentRequest, ID int64, UserID int64) (*dc.UpdateCommentResponse, error) {
	var err error

	q := fmt.Sprintf("%s %s %s %s %s %s", uqUpdateComment, uqWhere, uqFilterID, uqAnd, uqFilterUserID, uqReturningComment)

	query, args, err := ur.DBList.Backend.Read.In(q, data.Message, ID, UserID)
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

	var comment = dc.UpdateCommentResponse{}
	err = res.Scan(&comment.ID, &comment.Message, &comment.PhotoID, &comment.UserID, &comment.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &comment, nil
}

func (ur CommentDataRepo) DeleteComment(tx *sql.Tx, ID int64, UserID int64) error {
	var err error

	q := fmt.Sprintf("%s %s %s %s %s", uqDeleteComment, uqWhere, uqFilterID, uqAnd, uqFilterUserID)

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
