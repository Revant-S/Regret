package utils

import (
	"Regret/domain"
	"Regret/simulator/simdomain"
)

func FilterPerformance(carrier *simdomain.Carrier) *simdomain.CarrierResponse {
	return &simdomain.CarrierResponse{
		Carrier: domain.Carrier{
			Name:   carrier.Name,
			Cost:   carrier.Cost,
			Status: carrier.Status,
		},
	}
}

func FilterPerformanceFromList(carriers []simdomain.Carrier) []simdomain.CarrierResponse {
	var filteredResponse []simdomain.CarrierResponse
	for _, carrier := range carriers {
		filteredCarrier := FilterPerformance(&carrier)
		filteredResponse = append(filteredResponse, *filteredCarrier)
	}
	return filteredResponse
}
