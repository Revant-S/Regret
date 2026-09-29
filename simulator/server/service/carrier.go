package service

import (
	"Regret/simulator/server/repository"
	"Regret/simulator/server/utils"
	"Regret/simulator/simdomain"
	"errors"
)

type carrierService struct {
	carrierRepository repository.CarrierRepository
}

func (c *carrierService) GetAllCarriers() ([]simdomain.CarrierResponse, error) {
	carriers, err := c.carrierRepository.GetAll()
	if err != nil {
		return nil, err
	}
	filteredCarriers := utils.FilterPerformanceFromList(carriers)
	return filteredCarriers, nil
}

func (c *carrierService) GetCarrierDetail(Id string) (*simdomain.CarrierResponse, error) {
	if Id == "" {
		return nil, errors.New("invalid carrierId")
	}
	carrier, err := c.carrierRepository.GetDetails(Id)
	if err != nil {
		return nil, err
	}

	filteredCarrier := utils.FilterPerformance(carrier)
	return filteredCarrier, nil
}

func (c *carrierService) GetAvailableCarriers() ([]simdomain.CarrierResponse, error) {
	availCarriers, err := c.carrierRepository.GetAvailable()
	if err != nil {
		return nil, err
	}
	return utils.FilterPerformanceFromList(availCarriers), nil
}

type CarrierService interface {
	GetAllCarriers() ([]simdomain.CarrierResponse, error)
	GetCarrierDetail(Id string) (*simdomain.CarrierResponse, error)
	GetAvailableCarriers() ([]simdomain.CarrierResponse, error)
}

func NewCarrierService(carrierRepo repository.CarrierRepository) CarrierService {
	return &carrierService{
		carrierRepository: carrierRepo,
	}
}
