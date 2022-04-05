package user_auth

import (
	"fmt"
	"log"

	uadb "github.com/bhanupbalusu/custpreorderms/data/mongodb/user_auth"
	r "github.com/bhanupbalusu/custpreorderms/domain/application_interface/repo"
)

func NewDBConnection() r.UserAuthRepoInterface {
	fmt.Println("............ Starting New MongoDB Connection .............")
	repo, err := uadb.NewMongoRepository("mongodb://localhost:27017", "cust_ecomm_app_v001", 30)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(repo)
	return repo
}
