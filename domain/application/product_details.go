package application

import (
	"fmt"

	r "github.com/bhanupbalusu/custpreorderms/domain/application_interface/repo"
	s "github.com/bhanupbalusu/custpreorderms/domain/application_interface/service"
	m "github.com/bhanupbalusu/custpreorderms/domain/model/product_details"
)

type productDetailsService struct {
	pdri r.ProductDetailsRepoInterface
}

func NewProductDetailsService(pdri r.ProductDetailsRepoInterface) s.ProductDetailsServiceInterface {
	return &productDetailsService{pdri}
}

func (pds *productDetailsService) Get() (m.ProductDetailsModelList, error) {
	return pds.pdri.Get()
}

func (pds *productDetailsService) GetByID(id string) (*m.ProductDetailsModel, error) {
	fmt.Println("---- from inside domain.controller.product_service.GetByID -----")
	return pds.pdri.GetByID(id)
}

func (pds *productDetailsService) Create(pm *m.ProductDetailsModel) (string, error) {
	return pds.pdri.Create(pm)
}

func (pds *productDetailsService) Update(pm *m.ProductDetailsModel) error {
	return pds.pdri.Update(pm)
}

func (pds *productDetailsService) Delete(pid string) error {
	return pds.pdri.Delete(pid)
}
