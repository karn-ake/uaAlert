package repository

type Repository interface{
	FindAll() ([]string,error)
}