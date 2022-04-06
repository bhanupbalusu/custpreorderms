package application

import (
	r "github.com/bhanupbalusu/custpreorderms/domain/application_interface/repo"
	s "github.com/bhanupbalusu/custpreorderms/domain/application_interface/service"
	m "github.com/bhanupbalusu/custpreorderms/domain/model/deal_settings"
)

type dealSettingsService struct {
	dsri r.DealSettingsRepoInterface
}

func NewDealSettingsService(dsri r.DealSettingsRepoInterface) s.DealSettingsServiceInterface {
	return &dealSettingsService{dsri}
}

func (dss *dealSettingsService) Get() (*[]m.DealSettingsModel, error) {
	return dss.dsri.Get()
}

func (dss *dealSettingsService) GetByID(id string) (*m.DealSettingsModel, error) {
	return dss.dsri.GetByID(id)
}

func (dss *dealSettingsService) Create(pm *m.DealSettingsModel) (string, error) {
	return dss.dsri.Create(pm)
}

func (dss *dealSettingsService) Update(pm *m.DealSettingsModel) error {
	return dss.dsri.Update(pm)
}

func (dss *dealSettingsService) Delete(pid string) error {
	return dss.dsri.Delete(pid)
}
