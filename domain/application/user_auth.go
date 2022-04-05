package application

import (
	"fmt"

	r "github.com/bhanupbalusu/custpreorderms/domain/application_interface/repo"
	s "github.com/bhanupbalusu/custpreorderms/domain/application_interface/service"
	m "github.com/bhanupbalusu/custpreorderms/domain/model/user_auth"
)

type userAuthService struct {
	uari r.UserAuthRepoInterface
}

func NewUserAuthService(uari r.UserAuthRepoInterface) s.UserAuthServiceInterface {
	return &userAuthService{uari: uari}
}

func (uas *userAuthService) Get() (m.UserList, error) {
	return uas.uari.Get()
}

func (uas *userAuthService) GetByID(id string) (*m.User, error) {
	return uas.uari.GetByID(id)
}

func (uas *userAuthService) GetByEmail(email string) (*m.User, error) {
	fmt.Println("---------Application.Create before calling Mongodb.GetByEmail---------")
	return uas.uari.GetByEmail(email)
}

func (uas *userAuthService) Create(u *m.User) error {
	fmt.Println("---------Application.Create before calling Mongodb.Create---------")
	return uas.uari.Create(u)
}

func (uas *userAuthService) Update(u *m.User) error {
	return uas.uari.Update(u)
}

func (uas *userAuthService) Delete(id string) error {
	return uas.uari.Delete(id)
}
