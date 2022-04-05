package repo

import (
	model "github.com/bhanupbalusu/custpreorderms/domain/model/order_types"
)

type OrderTypesRepoInterface interface {
	Get() (model.OrderTypesModelList, error)
	GetByID(id string) (*model.OrderTypesModel, error)
	Create(otm *model.OrderTypesModel) (string, error)
	Update(otm *model.OrderTypesModel) error
	Delete(id string) error
}
