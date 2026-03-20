package custom_err

import "errors"

var ErrBuildingQuery = errors.New("Error while building query ")
var ErrUserIsExist = errors.New("User is exist ")
var ErrCreateUser = errors.New("Error while creating user ")

var ErrInvalidArguments = errors.New("Invalid arguments ")

var ErrParseTokenToString = errors.New("Error while parsing token string ")

var ErrNilSecret = errors.New("Error while getting secret ")
