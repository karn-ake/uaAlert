package repository

import (
	"github.com/stretchr/testify/mock"
)

type mongoRepositoryMock struct {
	mock.Mock
}

func NewMock() *mongoRepositoryMock {
	return &mongoRepositoryMock{}
}

func (m *mongoRepositoryMock) FindAll() ([]Client, error) {
	args := m.Called()
	return args.Get(0).([]Client),args.Error(1)
}

func (m *mongoRepositoryMock) FindbyClientName(cn string) (*Client, error) {
	return nil, nil
}

func (m *mongoRepositoryMock) IsClientNameAdded(cn string) (bool, error) {
	return true, nil
}

func (m *mongoRepositoryMock) Update() error {
	return nil
}

func (m *mongoRepositoryMock) DelAll() error {
	panic(0)
}
