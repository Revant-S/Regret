package routes

import (
	"Regret/simulator/server/controllers"
	"github.com/labstack/echo/v5"
)

type carrierRouter struct {
	carrierController controllers.CarrierController
}

type CarrierRouter interface {
	RegisterCarrierRoutes(e *echo.Echo)
}

func NewCarrierRouter(controller controllers.CarrierController) CarrierRouter {
	return &carrierRouter{
		carrierController: controller,
	}
}

func (cr *carrierRouter) RegisterCarrierRoutes(e *echo.Echo) {
	carrier := e.Group("/carrier")
	controller := cr.carrierController
	carrier.GET("/all", controller.GetAllCarriersHandler)
	carrier.GET("/all/available", controller.GetAvailableCarriersHandler)
	carrier.GET("/:Id", controller.GetCarrierDetailHandler)
}
