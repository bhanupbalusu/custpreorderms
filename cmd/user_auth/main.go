package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	c "github.com/bhanupbalusu/custpreorderms/api/controller"
	h "github.com/bhanupbalusu/custpreorderms/api/handler/user_auth"
	a "github.com/bhanupbalusu/custpreorderms/domain/application"
	i "github.com/bhanupbalusu/custpreorderms/pkg/initialize"
	db "github.com/bhanupbalusu/custpreorderms/pkg/initialize/user_auth"
)

func main() {
	conn := db.NewDBConnection()
	service := a.NewUserAuthService(conn)
	controller := c.NewUserAuthController(service)
	routes := h.NewAuthRoutesHandler(controller)

	app := i.NewFiberApp()
	routes.Install(app)

	errs := make(chan error, 2)
	go func() {
		fmt.Println("Listening on port :4001")
		errs <- app.Listen(":4001")
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()
	fmt.Printf("Terminated %s", <-errs)
}
