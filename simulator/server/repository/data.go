package repository

import "Regret/simulator/domain"

var DummyCarriers []domain.Carrier = []domain.Carrier{
	{
		Id:          "CARRIER-PREMIUM-1",
		Cost:        300,
		Reliability: 0.99,
		Status:      domain.CarrierAvailable,
	},
	{
		Id:          "CARRIER-PREMIUM-2",
		Cost:        220,
		Reliability: 0.96,
		Status:      domain.CarrierAvailable,
	},
	{
		Id:          "CARRIER-STANDARD-1",
		Cost:        150,
		Reliability: 0.90,
		Status:      domain.CarrierAvailable,
	},
	{
		Id:          "CARRIER-STANDARD-2",
		Cost:        120,
		Reliability: 0.85,
		Status:      domain.CarrierAvailable,
	},
	{
		Id:          "CARRIER-STANDARD-3",
		Cost:        95,
		Reliability: 0.80,
		Status:      domain.CarrierAvailable,
	},
	{
		Id:          "CARRIER-ECONOMY-1",
		Cost:        60,
		Reliability: 0.70,
		Status:      domain.CarrierAvailable,
	},
	{
		Id:          "CARRIER-ECONOMY-2",
		Cost:        40,
		Reliability: 0.55,
		Status:      domain.CarrierAvailable,
	},
	{
		Id:          "CARRIER-GAMBLE-1",
		Cost:        15,
		Reliability: 0.30,
		Status:      domain.CarrierAvailable,
	},
	{
		Id:          "CARRIER-GAMBLE-2",
		Cost:        5,
		Reliability: 0.10,
		Status:      domain.CarrierAvailable,
	},
	{
		Id:          "CARRIER-OFFLINE",
		Cost:        100,
		Reliability: 0.88,
		Status:      domain.CarrierUnAvailable,
	},
}
