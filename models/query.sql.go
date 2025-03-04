// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package models

import (
	"context"
	"database/sql"
	"time"
)

const createPhoto = `-- name: CreatePhoto :execresult
INSERT INTO photo (user_id, title, description, filename, date) VALUES (?, ?, ?, ?, ?)
`

type CreatePhotoParams struct {
	UserID      uint32
	Title       string
	Description sql.NullString
	Filename    string
	Date        time.Time
}

func (q *Queries) CreatePhoto(ctx context.Context, arg CreatePhotoParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createPhoto,
		arg.UserID,
		arg.Title,
		arg.Description,
		arg.Filename,
		arg.Date,
	)
}

const createUser = `-- name: CreateUser :execresult
INSERT INTO user (username, email, password) VALUES (?, ?, ?)
`

type CreateUserParams struct {
	Username string
	Email    string
	Password string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createUser, arg.Username, arg.Email, arg.Password)
}

const deletePhoto = `-- name: DeletePhoto :execresult
DELETE FROM photo WHERE photo_id = ?
`

func (q *Queries) DeletePhoto(ctx context.Context, photoID uint32) (sql.Result, error) {
	return q.db.ExecContext(ctx, deletePhoto, photoID)
}

const emailExists = `-- name: EmailExists :one
SELECT EXISTS(SELECT 1 FROM user WHERE email = ?)
`

func (q *Queries) EmailExists(ctx context.Context, email string) (bool, error) {
	row := q.db.QueryRowContext(ctx, emailExists, email)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const getPhoto = `-- name: GetPhoto :one
SELECT photo_id, user_id, title, description, filename, date FROM photo WHERE photo_id = ?
`

func (q *Queries) GetPhoto(ctx context.Context, photoID uint32) (Photo, error) {
	row := q.db.QueryRowContext(ctx, getPhoto, photoID)
	var i Photo
	err := row.Scan(
		&i.PhotoID,
		&i.UserID,
		&i.Title,
		&i.Description,
		&i.Filename,
		&i.Date,
	)
	return i, err
}

const getPhotos = `-- name: GetPhotos :many
SELECT photo_id, user_id, title, description, filename, date FROM photo
`

func (q *Queries) GetPhotos(ctx context.Context) ([]Photo, error) {
	rows, err := q.db.QueryContext(ctx, getPhotos)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Photo
	for rows.Next() {
		var i Photo
		if err := rows.Scan(
			&i.PhotoID,
			&i.UserID,
			&i.Title,
			&i.Description,
			&i.Filename,
			&i.Date,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT user_id, username, email, password, creted_at FROM user WHERE email = ?
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.CretedAt,
	)
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT user_id, username, email, password, creted_at FROM user WHERE user_id = ?
`

func (q *Queries) GetUserById(ctx context.Context, userID uint32) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserById, userID)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.CretedAt,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT user_id, username, email, password, creted_at FROM user WHERE username = ?
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByUsername, username)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.CretedAt,
	)
	return i, err
}

const getUsers = `-- name: GetUsers :many
SELECT user_id, username, email, password, creted_at FROM user
`

func (q *Queries) GetUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.UserID,
			&i.Username,
			&i.Email,
			&i.Password,
			&i.CretedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const usernameExists = `-- name: UsernameExists :one
SELECT EXISTS(SELECT 1 FROM user WHERE username = ?)
`

func (q *Queries) UsernameExists(ctx context.Context, username string) (bool, error) {
	row := q.db.QueryRowContext(ctx, usernameExists, username)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}
