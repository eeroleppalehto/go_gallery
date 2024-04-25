-- name: GetUsers :many
SELECT * FROM user;

-- name: GetUserByUsername :one
SELECT * FROM user WHERE username = ?;

-- name: GetUserById :one
SELECT * FROM user WHERE user_id = ?;

-- name: GetUserByEmail :one
SELECT * FROM user WHERE email = ?;

-- name: EmailExists :one
SELECT EXISTS(SELECT 1 FROM user WHERE email = ?);

-- name: UsernameExists :one
SELECT EXISTS(SELECT 1 FROM user WHERE username = ?);

-- name: CreateUser :execresult
INSERT INTO user (username, email, password) VALUES (?, ?, ?);

-- name: GetPhotos :many
SELECT * FROM photo;

-- name: CreatePhoto :execresult
INSERT INTO photo (user_id, title, description, filename, date) VALUES (?, ?, ?, ?, ?);

-- name: GetPhoto :one
SELECT * FROM photo WHERE photo_id = ?;
