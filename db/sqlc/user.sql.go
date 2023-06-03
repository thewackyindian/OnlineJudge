// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: user.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  name,
  username,
  email,
  password,
  dob
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING id, name, username, email, password, profileimg, motto, created_at, dob, rating, problem_solved, admin_id, is_setter
`

type CreateUserParams struct {
	Name     string    `json:"name"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Dob      time.Time `json:"dob"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Name,
		arg.Username,
		arg.Email,
		arg.Password,
		arg.Dob,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.Profileimg,
		&i.Motto,
		&i.CreatedAt,
		&i.Dob,
		&i.Rating,
		&i.ProblemSolved,
		&i.AdminID,
		&i.IsSetter,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE username = $1
`

func (q *Queries) DeleteUser(ctx context.Context, username string) error {
	_, err := q.db.ExecContext(ctx, deleteUser, username)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, name, username, email, password, profileimg, motto, created_at, dob, rating, problem_solved, admin_id, is_setter FROM users
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.Profileimg,
		&i.Motto,
		&i.CreatedAt,
		&i.Dob,
		&i.Rating,
		&i.ProblemSolved,
		&i.AdminID,
		&i.IsSetter,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, name, username, email, password, profileimg, motto, created_at, dob, rating, problem_solved, admin_id, is_setter FROM users
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
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Username,
			&i.Email,
			&i.Password,
			&i.Profileimg,
			&i.Motto,
			&i.CreatedAt,
			&i.Dob,
			&i.Rating,
			&i.ProblemSolved,
			&i.AdminID,
			&i.IsSetter,
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

const updateUser = `-- name: UpdateUser :one
UPDATE users
  set name = $2,
  email = $3,
  password = $4,
  profileimg = $5,
  motto = $6,
  dob = $7,
  is_setter = $8
WHERE username = $1
RETURNING id, name, username, email, password, profileimg, motto, created_at, dob, rating, problem_solved, admin_id, is_setter
`

type UpdateUserParams struct {
	Username   string         `json:"username"`
	Name       string         `json:"name"`
	Email      string         `json:"email"`
	Password   string         `json:"password"`
	Profileimg sql.NullString `json:"profileimg"`
	Motto      sql.NullString `json:"motto"`
	Dob        time.Time      `json:"dob"`
	IsSetter   bool           `json:"is_setter"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.Username,
		arg.Name,
		arg.Email,
		arg.Password,
		arg.Profileimg,
		arg.Motto,
		arg.Dob,
		arg.IsSetter,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.Profileimg,
		&i.Motto,
		&i.CreatedAt,
		&i.Dob,
		&i.Rating,
		&i.ProblemSolved,
		&i.AdminID,
		&i.IsSetter,
	)
	return i, err
}
