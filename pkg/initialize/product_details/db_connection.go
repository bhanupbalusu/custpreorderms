package product_details

import (
	"fmt"
	"log"

	pddb "github.com/bhanupbalusu/custpreorderms/data/mongodb/product_details"
	r "github.com/bhanupbalusu/custpreorderms/domain/application_interface/repo"
)

func NewDBConnection() r.ProductDetailsRepoInterface {
	fmt.Println("............ Starting New MongoDB Connection .............")
	repo, err := pddb.NewMongoRepository("mongodb://localhost:27017", "cust_ecomm_app_v001", 30)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(repo)
	return repo
}
