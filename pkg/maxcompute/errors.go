package maxcompute

import "errors"

var (
	ErrorMessageInvalidJSON            = errors.New("could not parse json")
	ErrorMessageInvalidEndpoint        = errors.New("invalid endpoint. Either empty or not set")
	ErrorMessageInvalidProjectName     = errors.New("invalid project name. Either empty or not set")
	ErrorMessageInvalidAccessKeyId     = errors.New("access key id is either empty or not set")
	ErrorMessageInvalidAccessKeySecret = errors.New("access key secret is either empty or not set")
)
