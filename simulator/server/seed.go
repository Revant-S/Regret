package server

import (
	"Regret/domain"
	"Regret/simulator/server/data"
	"Regret/simulator/server/models"
	"gorm.io/gorm"
)

func SeedDB(db *gorm.DB) {

	// seed Carriers to DB
	for _, carrier := range data.SeedCarriers {
		carry := &models.Carrier{
			Carrier: domain.Carrier{
				Cost:   carrier.Cost,
				Status: carrier.Status,
			},
			Reliability: carrier.Reliability,
			FailureRate: carrier.FailureRate,
		}
		db.Where(carrier).FirstOrCreate(carry)
	}
}
