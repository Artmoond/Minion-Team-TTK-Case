package custom_err

import "errors"

var ErrBuildingQuery = errors.New("Error while building query ")
var ErrUserIsExist = errors.New("User is exist ")
var ErrUserNotFound = errors.New("User not found ")
var ErrCreateUser = errors.New("Error while creating user ")
var ErrLoginUser = errors.New("Error while logining user ")
var ErrPasswordNotAddUp = errors.New("Error password doesn't add up ")

var ErrInvalidArguments = errors.New("Invalid arguments ")

var ErrParseTokenToString = errors.New("Error while parsing token string ")

var ErrNilSecret = errors.New("Error while getting secret ")
