package service

import (
	model "github.com/bhanupbalusu/custpreorderms/domain/model/preorder_metadata"
)

type PreOrderServiceInterface interface {
	Get() (model.PreOrderMetaDataModelList, error)
	GetByID(id string) (*model.PreOrderMetaDataModel, error)
	Create(pomd *model.PreOrderMetaDataModel) (string, error)
	Delete(poid string) error
}
