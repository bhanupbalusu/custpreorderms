package deal_settings

import (
	"fmt"
	"log"

	dsdb "github.com/bhanupbalusu/custpreorderms/data/mongodb/deal_settings"
	r "github.com/bhanupbalusu/custpreorderms/domain/application_interface/repo"
)

func NewDBConnection() r.DealSettingsRepoInterface {
	fmt.Println("............ Starting New MongoDB Connection .............")
	repo, err := dsdb.NewMongoRepository("mongodb://localhost:27017", "cust_ecomm_app_v001", 30)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(repo)
	return repo
}
