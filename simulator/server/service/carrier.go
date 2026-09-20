package service

import (
	"Regret/simulator/impl"
	"errors"
	"slices"
)

type CarrierService struct {
	AvailableCarriers []impl.CarrierResponse
}

func (c *CarrierService) GetAllCarriers() ([]impl.CarrierResponse, error) {
	return c.AvailableCarriers, nil
}

func (c *CarrierService) GetCarrierDetail(Id string) (*impl.CarrierResponse, error) {
	if Id == "" {
		return nil, errors.New("invalid carrierId")
	}
	carrierIndex := slices.IndexFunc(c.AvailableCarriers, func(carrier impl.CarrierResponse) bool {
		return carrier.Id == Id
	})
	if carrierIndex == -1 {
		return nil, errors.New("carrier not found")
	}
	return &(c.AvailableCarriers[carrierIndex]), nil
}

func (c *CarrierService) GetAvailableCarriers() ([]impl.CarrierResponse, error) {
	var availableCarriers []impl.CarrierResponse
	for _, carrier := range c.AvailableCarriers {
		if carrier.Status == impl.CarrierAvailable {
			availableCarriers = append(availableCarriers, carrier)
		}
	}
	return availableCarriers, nil
}

type CarrierServiceInterface interface {
	GetAllCarriers() ([]impl.CarrierResponse, error)
	GetCarrierDetail(Id string) (*impl.CarrierResponse, error)
	GetAvailableCarriers() ([]impl.CarrierResponse, error)
}

func NewCarrierService() CarrierServiceInterface {
	return &CarrierService{
		AvailableCarriers: []impl.CarrierResponse{
			{
				Id:     "CARRIER-PREMIUM-1",
				Cost:   300,
				Status: impl.CarrierAvailable,
			},
			{
				Id:     "CARRIER-PREMIUM-2",
				Cost:   220,
				Status: impl.CarrierAvailable,
			},
			{
				Id:     "CARRIER-STANDARD-1",
				Cost:   150,
				Status: impl.CarrierAvailable,
			},
			{
				Id:     "CARRIER-STANDARD-2",
				Cost:   120,
				Status: impl.CarrierAvailable,
			},
			{
				Id:     "CARRIER-STANDARD-3",
				Cost:   95,
				Status: impl.CarrierAvailable,
			},
			{
				Id:     "CARRIER-ECONOMY-1",
				Cost:   60,
				Status: impl.CarrierAvailable,
			},
			{
				Id:     "CARRIER-ECONOMY-2",
				Cost:   40,
				Status: impl.CarrierAvailable,
			},
			{
				Id:     "CARRIER-GAMBLE-1",
				Cost:   15,
				Status: impl.CarrierAvailable,
			},
			{
				Id:     "CARRIER-GAMBLE-2",
				Cost:   5,
				Status: impl.CarrierAvailable,
			},
			{
				Id:     "CARRIER-OFFLINE",
				Cost:   100,
				Status: impl.CarrierUnAvailable,
			},
		},
	}
}
