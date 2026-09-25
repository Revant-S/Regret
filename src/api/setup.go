package api

import (
	"Regret/src/api/handlers"
	"Regret/src/api/managers"
	routers "Regret/src/api/router"
	"Regret/src/api/schema"
	"Regret/src/api/storage"
	"github.com/labstack/echo/v5"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func setupDispatch(e *echo.Echo, DB *gorm.DB) {
	store := storage.NewCarrierStorage(DB)
	manage := managers.NewDispatchManager(store)
	handle := handlers.NewDispatchHandler(manage)
	router := routers.NewDispatchRouter(handle)

	router.RegisterDispatchRouter(e)
}

func SetUpDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("policy_store"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error Connecting To DB %v", err)
	}
	log.Println("Starting Migration...")
	err = db.AutoMigrate(&schema.Carrier{})
	if err != nil {
		log.Fatalf("Error In Migration %v", err)
	}
	log.Println("Migration SuccessFul")
	return db
}

func SetUp(e *echo.Echo) {
	db := SetUpDB()
	setupDispatch(e, db)
}
