package service

import (
	model "github.com/bhanupbalusu/custpreorderms/domain/model/order_types"
)

type OrderTypesServiceInterface interface {
	Get() (*[]model.OrderTypesModel, error)
	GetByID(id string) (*model.OrderTypesModel, error)
	Create(otm *model.OrderTypesModel) (string, error)
	Update(otm *model.OrderTypesModel) error
	Delete(id string) error
}
