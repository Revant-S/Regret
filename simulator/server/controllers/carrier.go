package controllers

import (
	"Regret/domain"
	"Regret/simulator/server/service"
	"Regret/simulator/simdomain"
	"github.com/labstack/echo/v5"
	"net/http"
)

type CarrierController struct {
	carrierService service.CarrierServiceInterface
}

func (cc *CarrierController) GetAllCarriersHandler(c *echo.Context) error {
	carriers, err := cc.carrierService.GetAllCarriers()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, carriers)
}

func (cc *CarrierController) GetCarrierDetailHandler(c *echo.Context) error {
	Id := c.Param("Id")
	if Id == "" {
		return c.JSON(http.StatusBadRequest, simdomain.CarrierResponse{
			Carrier: domain.Carrier{
				Id:     "",
				Cost:   0.0,
				Status: domain.CarrierInvalid,
			},
		})
	}

	carrier, err := cc.carrierService.GetCarrierDetail(Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, simdomain.CarrierResponse{})
	}
	return c.JSON(http.StatusOK, carrier)
}

func (cc *CarrierController) GetAvailableCarriersHandler(c *echo.Context) error {
	carriers, err := cc.carrierService.GetAvailableCarriers()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, carriers)
}

type CarrierControllerInterface interface {
	GetAllCarriersHandler(c *echo.Context) error
	GetCarrierDetailHandler(c *echo.Context) error
	GetAvailableCarriersHandler(c *echo.Context) error
}

func NewCarrierController(carrierServiceInterface service.CarrierServiceInterface) CarrierControllerInterface {
	return &CarrierController{
		carrierService: carrierServiceInterface,
	}
}
