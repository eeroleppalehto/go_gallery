-- name: GetUsers :many
SELECT * FROM user;

-- name: CreateUser :execresult
INSERT INTO user (username, email, password) VALUES (?, ?, ?);

-- name: GetPhotos :many
SELECT * FROM photo;

-- name: CreatePhoto :execresult
INSERT INTO photo (user_id, title, description, filename, date) VALUES (?, ?, ?, ?, ?);

-- name: GetPhoto :one
SELECT * FROM photo WHERE photo_id = ?;
