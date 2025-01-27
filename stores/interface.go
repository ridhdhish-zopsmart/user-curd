package stores

import (
	"layer/user/models"
)

type User interface {
	GetUserById(id int) (*models.User, error)
}
