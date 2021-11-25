package services

import "errors"

var (
	ErrGetLocalLogTime = errors.New("get local log time error")
	ErrParse           = errors.New("get parse time error")
	ErrGetAllTime      = errors.New("get all time error")
	ErrRevFile         = errors.New("get reverse file error")
	ErrOpen            = errors.New("get file open error")
)
