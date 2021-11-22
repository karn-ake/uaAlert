package services

import "time"

type Services interface {
	RevFile(fn string) (*[]string, error)
	GetLocalLogTime(cn string, fn string) (*string, error)
	GetAllTimes(cn string, lf string) (*AllTime, error)
	CheckValidate(dt time.Duration) bool
	CheckStatus(cn string, lf string) (*Customer, error)
}
