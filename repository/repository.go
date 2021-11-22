package repository

import "go.mongodb.org/mongo-driver/mongo"

type Repository interface {
	FindAll() ([]Client, error)
	FindbyClientName(cn string) (*Client, error)
	IsClientNameAdded(cn string) (bool, error)
	Update() error
	DelAll() (*mongo.DeleteResult, error)
}
