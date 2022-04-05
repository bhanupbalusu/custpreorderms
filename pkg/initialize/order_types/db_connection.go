package order_types

import (
	"fmt"
	"log"

	otdb "github.com/bhanupbalusu/custpreorderms/data/mongodb/order_types"
	r "github.com/bhanupbalusu/custpreorderms/domain/application_interface/repo"
)

func NewDBConnection() r.OrderTypesRepoInterface {
	fmt.Println("............ Starting New MongoDB Connection .............")
	repo, err := otdb.NewMongoRepository("mongodb://localhost:27017", "cust_ecomm_app_v001", 30)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(repo)
	return repo
}
