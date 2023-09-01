// Code generated by sqlc. DO NOT EDIT.

package mysql

import ()

type Skin struct {
	ID       int64  `db:"id"`
	UserID   int64  `db:"user_id"`
	SkinHash string `db:"skin_hash"`
	Type     string `db:"type"`
	Variant  string `db:"variant"`
}

type User struct {
	ID       int64  `db:"id"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Salt     string `db:"salt"`
	State    int32  `db:"state"`
	RegTime  int64  `db:"reg_time"`
}

type UserProfile struct {
	UserID int64  `db:"user_id"`
	Name   string `db:"name"`
	Uuid   string `db:"uuid"`
}

type UserSkin struct {
	UserID int64 `db:"user_id"`
	SkinID int64 `db:"skin_id"`
}

type UserToken struct {
	UserID  int64 `db:"user_id"`
	TokenID int32 `db:"token_id"`
}
