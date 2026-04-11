package service

import "errors"

var (
	ErrInvalidEmailFormat        = errors.New("invalid email format")
	ErrInvalidRepoFormat         = errors.New("invalid repository format, has to be \"owner/repo\"")
	ErrRepoNotFound              = errors.New("repository not found")
	ErrSubscriptionAlreadyExists = errors.New("subscription with such email and repo pair already exists")
	ErrTooMuchRequests           = errors.New("request limit hit")
)
