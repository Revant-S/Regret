package server

import (
	"Regret/simulator/server/controllers"
	"Regret/simulator/server/models"
	"Regret/simulator/server/repository"
	"Regret/simulator/server/routes"
	"Regret/simulator/server/service"
	"github.com/labstack/echo/v5"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func SetUpDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("sim_storage"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error Connecting To DB %v", err)
	}
	log.Println("Starting Migration...")
	err = db.AutoMigrate(&models.Carrier{})
	if err != nil {
		log.Fatalf("Error In Migration %v", err)
	}
	log.Println("Migration SuccessFul")
	return db
}

func setUpCarrier(e *echo.Echo, DB *gorm.DB) {
	repo := repository.NewCarrierRepository(DB)
	svc := service.NewCarrierService(repo)
	control := controllers.NewCarrierController(svc)
	rtr := routes.NewCarrierRouter(control)
	rtr.RegisterCarrierRoutes(e)
}

func SetUp(e *echo.Echo) {
	db := SetUpDB()
	setUpCarrier(e, db)
	SeedDB(db)

}
