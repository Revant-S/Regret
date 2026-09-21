package repository

import (
	"Regret/domain"
	"Regret/simulator/simdomain"
)

var DummyCarriers = []simdomain.Carrier{
	{
		Carrier: domain.Carrier{
			Name:   "CARRIER-PREMIUM-1",
			Cost:   300,
			Status: domain.CarrierAvailable,
		},
		Reliability: 0.99,
	},
	{
		Carrier: domain.Carrier{
			Name:   "CARRIER-PREMIUM-2",
			Cost:   220,
			Status: domain.CarrierAvailable,
		},
		Reliability: 0.96,
	},
	{
		Carrier: domain.Carrier{
			Name:   "CARRIER-STANDARD-1",
			Cost:   150,
			Status: domain.CarrierAvailable,
		},
		Reliability: 0.90,
	},
	{
		Carrier: domain.Carrier{
			Name:   "CARRIER-STANDARD-2",
			Cost:   120,
			Status: domain.CarrierAvailable,
		},
		Reliability: 0.85,
	},
	{
		Carrier: domain.Carrier{
			Name:   "CARRIER-STANDARD-3",
			Cost:   95,
			Status: domain.CarrierAvailable,
		},
		Reliability: 0.80,
	},
	{
		Carrier: domain.Carrier{
			Name:   "CARRIER-ECONOMY-1",
			Cost:   60,
			Status: domain.CarrierAvailable,
		},
		Reliability: 0.70,
	},
	{
		Carrier: domain.Carrier{
			Name:   "CARRIER-ECONOMY-2",
			Cost:   40,
			Status: domain.CarrierAvailable,
		},
		Reliability: 0.55,
	},
	{
		Carrier: domain.Carrier{
			Name:   "CARRIER-GAMBLE-1",
			Cost:   15,
			Status: domain.CarrierAvailable,
		},
		Reliability: 0.30,
	},
	{
		Carrier: domain.Carrier{
			Name:   "CARRIER-GAMBLE-2",
			Cost:   5,
			Status: domain.CarrierAvailable,
		},
		Reliability: 0.10,
	},
	{
		Carrier: domain.Carrier{
			Name:   "CARRIER-OFFLINE",
			Cost:   100,
			Status: domain.CarrierUnAvailable,
		},
		Reliability: 0.88,
	},
}
