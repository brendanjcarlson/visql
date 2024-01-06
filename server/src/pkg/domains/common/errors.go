package common

import "errors"

var (
	ErrNotYetImplemented = errors.New("not yet implemented")

	ErrFullNameRequired = errors.New("full name is required")

	ErrEmailRequired     = errors.New("email is required")
	ErrEmailAlreadyInUse = errors.New("email is already in use")

	ErrPasswordRequired = errors.New("password is required")
	ErrPasswordTooShort = errors.New("password must be at least 8 characters")
)

var (
	MsgNotYetImplemented = "not yet implemented"

	MsgInvalidRequestBody = "invalid request body"
	MsgInternalServerErr  = "internal server error"
)

func StringifyErrs(err ...error) []string {
	var errMsgs []string
	for _, err := range err {
		errMsgs = append(errMsgs, err.Error())
	}
	return errMsgs
}
