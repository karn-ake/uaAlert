package repository

type Repository interface {
	FindAll() ([]Client, error)
	FindbyClientName(cn string) (*Client, error)
	IsClientNameAdded(cn string) (bool, error)
	Update() error
	DelAll() error
}
