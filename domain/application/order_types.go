package application

import (
	r "github.com/bhanupbalusu/custpreorderms/domain/application_interface/repo"
	s "github.com/bhanupbalusu/custpreorderms/domain/application_interface/service"
	m "github.com/bhanupbalusu/custpreorderms/domain/model/order_types"
)

type orderTypesService struct {
	otri r.OrderTypesRepoInterface
}

func NewOrderTypesService(otri r.OrderTypesRepoInterface) s.OrderTypesServiceInterface {
	return &orderTypesService{otri}
}

func (ots *orderTypesService) Get() (m.OrderTypesModelList, error) {
	return ots.otri.Get()
}

func (ots *orderTypesService) GetByID(id string) (*m.OrderTypesModel, error) {
	return ots.otri.GetByID(id)
}

func (ots *orderTypesService) Create(otm *m.OrderTypesModel) (string, error) {
	return ots.otri.Create(otm)
}

func (ots *orderTypesService) Update(otm *m.OrderTypesModel) error {
	return ots.otri.Update(otm)
}

func (ots *orderTypesService) Delete(id string) error {
	return ots.otri.Delete(id)
}
