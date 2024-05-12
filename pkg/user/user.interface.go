package user

import (
	"invoiceApi/models"
)

type User interface {
	GetUser(id int) (*models.User, error)
	GetUsers() ([]models.User, error)
	UpdateUser(id int) (*models.User, error)
	InsertUser(user *models.User) error
}
