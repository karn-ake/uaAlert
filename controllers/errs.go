package controllers

import "errors"

var (
	errFindClient = errors.New("cannot find client by name")
	errCheckstatus = errors.New("cannot check log status")
)