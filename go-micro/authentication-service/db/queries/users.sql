-- name: CreateUser :one
INSERT INTO users (
    email,
    hased_password,
    active
) VALUES ($1, $2) RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;