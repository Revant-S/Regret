package utils

import "Regret/simulator/domain"

func FilterPerformance(carrier *domain.Carrier) *domain.CarrierResponse {
	return &domain.CarrierResponse{
		Id:     carrier.Id,
		Cost:   carrier.Cost,
		Status: carrier.Status,
	}
}

func FilterPerformanceFromList(carriers []domain.Carrier) []domain.CarrierResponse {
	var filteredResponse []domain.CarrierResponse
	for _, carrier := range carriers {
		filteredCarrier := FilterPerformance(&carrier)
		filteredResponse = append(filteredResponse, *filteredCarrier)
	}
	return filteredResponse
}
