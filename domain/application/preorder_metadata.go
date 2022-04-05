package application

import (
	r "github.com/bhanupbalusu/custpreorderms/domain/application_interface/repo"
	s "github.com/bhanupbalusu/custpreorderms/domain/application_interface/service"
	m "github.com/bhanupbalusu/custpreorderms/domain/model/preorder_metadata"
)

type preOrderService struct {
	pori r.PreOrderRepoInterface
}

func NewPreOrderService(pori r.PreOrderRepoInterface) s.PreOrderServiceInterface {
	return &preOrderService{pori}
}

func (pos *preOrderService) Get() (m.PreOrderMetaDataModelList, error) {
	return pos.pori.Get()
}

func (pos *preOrderService) GetByID(id string) (*m.PreOrderMetaDataModel, error) {
	return pos.pori.GetByID(id)
}

func (pos *preOrderService) Create(pm *m.PreOrderMetaDataModel) (string, error) {
	return pos.pori.Create(pm)
}

func (pos *preOrderService) Delete(id string) error {
	return pos.pori.Delete(id)
}
