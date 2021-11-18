package repository

import "go.mongodb.org/mongo-driver/mongo"

type Repository interface {
	FindAll() ([]Client, error)
	FindbyClientName(cn string) ([]Client, error)
	Update() error
	DelAll() (*mongo.DeleteResult, error)
}
