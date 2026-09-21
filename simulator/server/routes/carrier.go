package routes

import (
	"Regret/simulator/server/controllers"
	"github.com/labstack/echo/v5"
)

type CarrierRouter struct {
	carrierController controllers.CarrierControllerInterface
}

type CarrierRouterInterface interface {
	RegisterCarrierRoutes(e *echo.Echo)
}

func NewCarrierRouter(controller controllers.CarrierControllerInterface) CarrierRouterInterface {
	return &CarrierRouter{
		carrierController: controller,
	}
}

func (cr *CarrierRouter) RegisterCarrierRoutes(e *echo.Echo) {
	carrier := e.Group("/carrier")
	controller := cr.carrierController
	carrier.GET("/all", controller.GetAllCarriersHandler)
	carrier.GET("/all/available", controller.GetAvailableCarriersHandler)
	carrier.GET("/:Id", controller.GetCarrierDetailHandler)
}
