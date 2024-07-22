package constant

import (
	"errors"
)

const (
	CodeSuccess           = 200
	CodeErrBadRequest     = 400
	CodeErrUnauthorized   = 401
	CodeErrNotFound       = 404
	CodeAuth              = 406
	CodeGlobalLoading     = 407
	CodeErrInternalServer = 500

	CodeErrIP           = 310
	CodeErrDomain       = 311
	CodeErrEntrance     = 312
	CodePasswordExpired = 313

	CodeErrXpack = 410
)

// internal
var (
	ErrCaptchaCode     = errors.New("ErrCaptchaCode")
	ErrAuth            = errors.New("ErrAuth")
	ErrRecordExist     = errors.New("ErrRecordExist")
	ErrRecordNotFound  = errors.New("ErrRecordNotFound")
	ErrTransform       = errors.New("ErrTransform")
	ErrInitialPassword = errors.New("ErrInitialPassword")
	ErrInvalidParams   = errors.New("ErrInvalidParams")

	ErrTokenParse = errors.New("ErrTokenParse")
	ErrPortInUsed = "ErrPortInUsed"
	ErrCmdTimeout = "ErrCmdTimeout"
)

// api
var (
	ErrTypeInternalServer  = "ErrInternalServer"
	ErrTypeInvalidParams   = "ErrInvalidParams"
	ErrTypeNotLogin        = "ErrNotLogin"
	ErrTypePasswordExpired = "ErrPasswordExpired"
	ErrDemoEnvironment     = "ErrDemoEnvironment"
	ErrEntrance            = "ErrEntrance"
)
