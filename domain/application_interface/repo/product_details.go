package repo

import (
	model "github.com/bhanupbalusu/custpreorderms/domain/model/product_details"
)

type ProductDetailsRepoInterface interface {
	Get() (*[]model.ProductDetailsModel, error)
	GetByID(id string) (*model.ProductDetailsModel, error)
	Create(pm *model.ProductDetailsModel) (string, error)
	Update(pm *model.ProductDetailsModel) error
	Delete(pid string) error
}
