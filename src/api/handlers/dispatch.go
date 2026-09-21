package handlers

import (
	"Regret/src/api/managers"
	"github.com/labstack/echo/v5"
)

type DispatchHandler struct {
	dispatchManager managers.DispatchManagerInterface
}

type DispatchHandlerInterface interface {
	GetChannel(c *echo.Context) error
}

func NewDispatchHandler(manager managers.DispatchManagerInterface) DispatchHandlerInterface {
	return &DispatchHandler{
		dispatchManager: manager,
	}
}

func (dh *DispatchHandler) GetChannel(c *echo.Context) error {
	return nil
}
