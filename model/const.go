package model

type APIStatus int

const (
	ErrUnknown APIStatus = iota - 1
	OK
	ErrInput
	ErrService
	ErrExistUser
)
