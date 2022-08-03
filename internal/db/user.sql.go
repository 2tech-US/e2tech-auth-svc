// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: user.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  phone,
  password,
  role
) VALUES (
  $1, $2, $3
)
RETURNING id, phone, password, role, verified, created_at
`

type CreateUserParams struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Phone, arg.Password, arg.Role)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Phone,
		&i.Password,
		&i.Role,
		&i.Verified,
		&i.CreatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, phone, password, role, verified, created_at FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Phone,
		&i.Password,
		&i.Role,
		&i.Verified,
		&i.CreatedAt,
	)
	return i, err
}

const getUserByPhone = `-- name: GetUserByPhone :one
SELECT id, phone, password, role, verified, created_at FROM users
WHERE phone = $1 LIMIT 1
`

func (q *Queries) GetUserByPhone(ctx context.Context, phone string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByPhone, phone)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Phone,
		&i.Password,
		&i.Role,
		&i.Verified,
		&i.CreatedAt,
	)
	return i, err
}

const getUserForUpdate = `-- name: GetUserForUpdate :one
SELECT id, phone, password, role, verified, created_at FROM users
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE
`

func (q *Queries) GetUserForUpdate(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserForUpdate, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Phone,
		&i.Password,
		&i.Role,
		&i.Verified,
		&i.CreatedAt,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, phone, password, role, verified, created_at FROM users
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListUsersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Phone,
			&i.Password,
			&i.Role,
			&i.Verified,
			&i.CreatedAt,
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

const updatePassword = `-- name: UpdatePassword :one

UPDATE users
SET password = $2
WHERE id = $1
RETURNING id, phone, password, role, verified, created_at
`

type UpdatePasswordParams struct {
	ID       int64  `json:"id"`
	Password string `json:"password"`
}

// pagination: offset: skip many rows
func (q *Queries) UpdatePassword(ctx context.Context, arg UpdatePasswordParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updatePassword, arg.ID, arg.Password)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Phone,
		&i.Password,
		&i.Role,
		&i.Verified,
		&i.CreatedAt,
	)
	return i, err
}

const verify = `-- name: Verify :one
UPDATE users
SET verified = true
WHERE phone = $1
RETURNING id, phone, password, role, verified, created_at
`

func (q *Queries) Verify(ctx context.Context, phone string) (User, error) {
	row := q.db.QueryRowContext(ctx, verify, phone)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Phone,
		&i.Password,
		&i.Role,
		&i.Verified,
		&i.CreatedAt,
	)
	return i, err
}
