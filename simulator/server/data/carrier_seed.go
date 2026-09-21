package data

import (
	"Regret/domain"
	"Regret/simulator/simdomain"
)

var SeedCarriers = []simdomain.Carrier{
	{
		Carrier: domain.Carrier{
			Name:   "CARRIER-ENTERPRISE-X1",
			Cost:   1500.00,
			Status: domain.CarrierAvailable,
		},
		Reliability: 0.9999,
		FailureRate: 0.0001,
	},
	{
		Carrier: domain.Carrier{
			Name:   "CARRIER-PREMIUM-A",
			Cost:   500.00,
			Status: domain.CarrierAvailable,
		},
		Reliability: 0.985,
		FailureRate: 0.015,
	},
	{
		Carrier: domain.Carrier{
			Name:   "CARRIER-STANDARD-1",
			Cost:   250.00,
			Status: domain.CarrierAvailable,
		},
		Reliability: 0.95,
		FailureRate: 0.05,
	},
	{
		Carrier: domain.Carrier{
			Name:   "CARRIER-STANDARD-2",
			Cost:   210.00,
			Status: domain.CarrierAvailable,
		},
		Reliability: 0.92,
		FailureRate: 0.08,
	},
	{
		Carrier: domain.Carrier{
			Name:   "CARRIER-ECONOMY-A",
			Cost:   100.00,
			Status: domain.CarrierAvailable,
		},
		Reliability: 0.85,
		FailureRate: 0.15,
	},
	{
		Carrier: domain.Carrier{
			Name:   "CARRIER-BUDGET-1",
			Cost:   45.00,
			Status: domain.CarrierAvailable,
		},
		Reliability: 0.70,
		FailureRate: 0.30,
	},
	{
		Carrier: domain.Carrier{
			Name:   "CARRIER-GAMBLE-X",
			Cost:   10.00,
			Status: domain.CarrierAvailable,
		},
		Reliability: 0.35,
		FailureRate: 0.65,
	},
	{
		Carrier: domain.Carrier{
			Name:   "CARRIER-PREMIUM-MAINTENANCE",
			Cost:   480.00,
			Status: domain.CarrierUnAvailable,
		},
		Reliability: 0.99,
		FailureRate: 0.01,
	},
	{
		Carrier: domain.Carrier{
			Name:   "CARRIER-ECONOMY-OFFLINE",
			Cost:   95.00,
			Status: domain.CarrierUnAvailable,
		},
		Reliability: 0.80,
		FailureRate: 0.20,
	},
	{
		Carrier: domain.Carrier{
			Name:   "CARRIER-DEPRECATED-OLD",
			Cost:   0.00,
			Status: domain.CarrierInvalid,
		},
		Reliability: 0.00,
		FailureRate: 1.00,
	},
}
