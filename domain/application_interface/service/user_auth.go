package service

import (
	"github.com/bhanupbalusu/custpreorderms/domain/model"
)

type UserAuthServiceInterface interface {
	Get() (model.UserList, error)
	GetByID(id string) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	Create(u *model.User) error
	Update(u *model.User) error
	Delete(id string) error
}
