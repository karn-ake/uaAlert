package services

type Services interface {
	RevFile(fn string) (*[]string, error)
	GetLocalLogTime(fn string) (*string, error)
	GetAllTimes(lf string) (*AllTime, error)
}
