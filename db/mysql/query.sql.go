// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package mysql

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :execresult
REPLACE INTO user (
  id,
  email,
  password,
  salt,
  state,
  reg_time
)
VALUES
  (?, ?, ?, ?, ?, ?)
`

type CreateUserParams struct {
	ID       int64  `db:"id"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Salt     string `db:"salt"`
	State    int32  `db:"state"`
	RegTime  int64  `db:"reg_time"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (sql.Result, error) {
	return q.exec(ctx, q.createUserStmt, createUser,
		arg.ID,
		arg.Email,
		arg.Password,
		arg.Salt,
		arg.State,
		arg.RegTime,
	)
}

const createUserProfile = `-- name: CreateUserProfile :execresult
REPLACE INTO ` + "`" + `user_profile` + "`" + ` (` + "`" + `user_id` + "`" + `, ` + "`" + `name` + "`" + `, ` + "`" + `uuid` + "`" + `)
VALUES
  (?, ?, ?)
`

type CreateUserProfileParams struct {
	UserID int64  `db:"user_id"`
	Name   string `db:"name"`
	Uuid   string `db:"uuid"`
}

func (q *Queries) CreateUserProfile(ctx context.Context, arg CreateUserProfileParams) (sql.Result, error) {
	return q.exec(ctx, q.createUserProfileStmt, createUserProfile, arg.UserID, arg.Name, arg.Uuid)
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM
  user
WHERE
  id = ?
`

func (q *Queries) DeleteUser(ctx context.Context, id int64) error {
	_, err := q.exec(ctx, q.deleteUserStmt, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT
  id, email, password, salt, state, reg_time
FROM
  user
WHERE
  id = ?
LIMIT
  1
`

func (q *Queries) GetUser(ctx context.Context, id int64) (User, error) {
	row := q.queryRow(ctx, q.getUserStmt, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.Salt,
		&i.State,
		&i.RegTime,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT
  id, email, password, salt, state, reg_time
FROM
  user
WHERE
  email = ?
LIMIT
  1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.queryRow(ctx, q.getUserByEmailStmt, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.Salt,
		&i.State,
		&i.RegTime,
	)
	return i, err
}

const listUser = `-- name: ListUser :many
SELECT
  id, email, password, salt, state, reg_time
FROM
  user
ORDER BY
  reg_time
`

func (q *Queries) ListUser(ctx context.Context) ([]User, error) {
	rows, err := q.query(ctx, q.listUserStmt, listUser)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.Password,
			&i.Salt,
			&i.State,
			&i.RegTime,
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
