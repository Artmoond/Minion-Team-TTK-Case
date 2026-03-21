package custom_err

import "errors"

var ErrBuildingQuery = errors.New("Error while building query ")
var ErrGetAllUsers = errors.New("Error while getting all users ")

var ErrGetClaims = errors.New("Error while getting claims ")
var ErrTokenInvalid = errors.New("Error while validating token ")

var ErrEmptyToken = errors.New("Error token is empty ")
var ErrNilSecret = errors.New("Error while getting secret ")

var ErrNotHaveRightRole = errors.New("Error not have right role ")
