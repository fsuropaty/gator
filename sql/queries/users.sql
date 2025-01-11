-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES ( 
	$1,
	$2,
	$3,
	$4
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE name = $1;

-- name: DelUsersTable :exec
DELETE FROM users;

-- name: GetUsers :many
SELECT id, name FROM users ORDER BY name;
