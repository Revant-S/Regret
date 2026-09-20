package server

import (
	"Regret/simulator/server/controllers"
	"Regret/simulator/server/repository"
	"Regret/simulator/server/routes"
	"Regret/simulator/server/service"
	"github.com/labstack/echo/v5"
)

func setUpCarrier(e *echo.Echo) {
	repo := repository.NewCarrierRepository()
	svc := service.NewCarrierService(repo)
	control := controllers.NewCarrierController(svc)
	rtr := routes.NewCarrierRouter(control)
	rtr.RegisterControllerRouters(e)
}

func SetUp(e *echo.Echo) {
	setUpCarrier(e)
}
