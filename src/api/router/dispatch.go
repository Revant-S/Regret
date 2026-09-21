package router

import (
	"Regret/src/api/handlers"
	"github.com/labstack/echo/v5"
)

type DispatchRouter struct {
	handler handlers.DispatchHandlerInterface
}

type DispatchRouterInterface interface {
	RegisterDispatchRouter(e *echo.Echo)
}

func NewDispatchRouter(dispatchHandler handlers.DispatchHandlerInterface) DispatchRouterInterface {
	return &DispatchRouter{
		handler: dispatchHandler,
	}
}

func (dr *DispatchRouter) RegisterDispatchRouter(e *echo.Echo) {
}
