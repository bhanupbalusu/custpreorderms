package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	c "github.com/bhanupbalusu/custpreorderms/api/controller"
	a "github.com/bhanupbalusu/custpreorderms/domain/application"
	i "github.com/bhanupbalusu/custpreorderms/pkg/initialize"

	dsh "github.com/bhanupbalusu/custpreorderms/api/handler/deal_settings"
	oth "github.com/bhanupbalusu/custpreorderms/api/handler/order_types"
	poh "github.com/bhanupbalusu/custpreorderms/api/handler/preorder_metadata"
	pdh "github.com/bhanupbalusu/custpreorderms/api/handler/product_details"
	uah "github.com/bhanupbalusu/custpreorderms/api/handler/user_auth"

	dsdb "github.com/bhanupbalusu/custpreorderms/pkg/initialize/deal_settings"
	otdb "github.com/bhanupbalusu/custpreorderms/pkg/initialize/order_types"
	podb "github.com/bhanupbalusu/custpreorderms/pkg/initialize/preorder_metadata"
	pddb "github.com/bhanupbalusu/custpreorderms/pkg/initialize/product_details"
	uadb "github.com/bhanupbalusu/custpreorderms/pkg/initialize/user_auth"
)

func main() {
	uaconn := uadb.NewDBConnection()
	uaservice := a.NewUserAuthService(uaconn)
	uacontroller := c.NewUserAuthController(uaservice)
	uaroutes := uah.NewAuthRoutesHandler(uacontroller)

	poconn := podb.NewDBConnection()
	poservice := a.NewPreOrderService(poconn)
	pocontroller := c.NewPreOrderController(poservice)
	poroutes := poh.NewPreOrderRoutesHandler(pocontroller)

	pdconn := pddb.NewDBConnection()
	pdservice := a.NewProductDetailsService(pdconn)
	pdcontroller := c.NewProductDetailsController(pdservice)
	pdroutes := pdh.NewProductDetailsRoutesHandler(pdcontroller)

	otconn := otdb.NewDBConnection()
	otservice := a.NewOrderTypesService(otconn)
	otcontroller := c.NewOrderTypesController(otservice)
	otroutes := oth.NewOrderTypesRoutesHandler(otcontroller)

	dsconn := dsdb.NewDBConnection()
	dsservice := a.NewDealSettingsService(dsconn)
	dscontroller := c.NewDealSettingsController(dsservice)
	dsroutes := dsh.NewDealSettingsRoutesHandler(dscontroller)

	app := i.NewFiberApp()

	uaroutes.Install(app)
	poroutes.Install(app)
	pdroutes.Install(app)
	otroutes.Install(app)
	dsroutes.Install(app)

	errs := make(chan error, 2)
	go func() {
		fmt.Println("Listening on port :4000")
		errs <- app.Listen(":4000")
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()
	fmt.Printf("Terminated %s", <-errs)
}
