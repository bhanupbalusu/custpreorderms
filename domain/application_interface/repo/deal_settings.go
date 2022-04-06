package repo

import (
	model "github.com/bhanupbalusu/custpreorderms/domain/model/deal_settings"
)

type DealSettingsRepoInterface interface {
	Get() (*[]model.DealSettingsModel, error)
	GetByID(id string) (*model.DealSettingsModel, error)
	Create(pm *model.DealSettingsModel) (string, error)
	Update(pm *model.DealSettingsModel) error
	Delete(pid string) error
}
