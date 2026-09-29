package handlers

import (
	"Regret/src/api/managers"
	"github.com/labstack/echo/v5"
)

type dispatchHandler struct {
	dispatchManager managers.DispatchManager
}

type DispatchHandler interface {
	GetChannel(c *echo.Context) error
}

func NewDispatchHandler(manager managers.DispatchManager) DispatchHandler {
	return &dispatchHandler{
		dispatchManager: manager,
	}
}

func (dh *dispatchHandler) GetChannel(c *echo.Context) error {
	return nil
}
