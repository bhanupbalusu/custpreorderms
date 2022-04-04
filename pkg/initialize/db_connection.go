package initialize

import (
	"fmt"
	"log"

	db "github.com/bhanupbalusu/custpreorderms/data/mongodb"
	r "github.com/bhanupbalusu/custpreorderms/domain/application_interface/repo"
)

func NewDBConnection() r.UserAuthRepoInterface {
	fmt.Println("............ Starting New MongoDB Connection .............")
	repo, err := db.NewMongoRepository("mongodb://localhost:27017", "ecomm-app-v001", 30)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(repo)
	return repo
}
