package api

import (
	"Regret/src/api/handlers"
	"Regret/src/api/managers"
	router2 "Regret/src/api/router"
	"Regret/src/api/storage"
	"github.com/labstack/echo/v5"
)

func setupDispatch(e *echo.Echo) {
	store := storage.NewCarrierStorage()
	manage := managers.NewDispatchManager(store)
	handle := handlers.NewDispatchHandler(manage)
	router := router2.NewDispatchRouter(handle)

	router.RegisterDispatchRouter(e)
}

func SetUp(e *echo.Echo) {
	setupDispatch(e)
}
