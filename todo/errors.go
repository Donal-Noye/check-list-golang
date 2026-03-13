package todo

import "errors"

var ErrTaskAlreadyExist = errors.New("task already exists")
var ErrTaskNotFound = errors.New("task not found")
