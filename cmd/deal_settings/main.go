package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	c "github.com/bhanupbalusu/custpreorderms/api/controller"
	h "github.com/bhanupbalusu/custpreorderms/api/handler/deal_settings"
	a "github.com/bhanupbalusu/custpreorderms/domain/application"
	i "github.com/bhanupbalusu/custpreorderms/pkg/initialize"
	db "github.com/bhanupbalusu/custpreorderms/pkg/initialize/deal_settings"
)

func main() {
	conn := db.NewDBConnection()
	service := a.NewDealSettingsService(conn)
	controller := c.NewDealSettingsController(service)
	routes := h.NewDealSettingsRoutesHandler(controller)

	app := i.NewFiberApp()
	routes.Install(app)

	errs := make(chan error, 2)
	go func() {
		fmt.Println("Listening on port :4003")
		errs <- app.Listen(":4004")
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()
	fmt.Printf("Terminated %s", <-errs)
}
