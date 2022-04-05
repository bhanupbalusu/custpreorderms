package service

import (
	model "github.com/bhanupbalusu/custpreorderms/domain/model/product_details"
)

type ProductDetailsServiceInterface interface {
	Get() (model.ProductDetailsModelList, error)
	GetByID(id string) (*model.ProductDetailsModel, error)
	Create(pm *model.ProductDetailsModel) (string, error)
	Update(pm *model.ProductDetailsModel) error
	Delete(pid string) error
}
