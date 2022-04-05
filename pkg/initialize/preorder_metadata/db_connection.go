package preorder_metadata

import (
	"fmt"
	"log"

	podb "github.com/bhanupbalusu/custpreorderms/data/mongodb/preorder_metadata"
	r "github.com/bhanupbalusu/custpreorderms/domain/application_interface/repo"
)

func NewDBConnection() r.PreOrderRepoInterface {
	fmt.Println("............ Starting New MongoDB Connection .............")
	repo, err := podb.NewMongoRepository("mongodb://localhost:27017", "cust_ecomm_app_v001", 30)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(repo)
	return repo
}
