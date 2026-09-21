package service

import (
	"Regret/simulator/server/repository"
	"Regret/simulator/server/utils"
	"Regret/simulator/simdomain"
	"errors"
)

type CarrierService struct {
	carrierRepository repository.CarrierRepositoryInterface
}

func (c *CarrierService) GetAllCarriers() ([]simdomain.CarrierResponse, error) {
	carriers, err := c.carrierRepository.GetAll()
	if err != nil {
		return nil, err
	}
	filteredCarriers := utils.FilterPerformanceFromList(carriers)
	return filteredCarriers, nil
}

func (c *CarrierService) GetCarrierDetail(Id string) (*simdomain.CarrierResponse, error) {
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

func (c *CarrierService) GetAvailableCarriers() ([]simdomain.CarrierResponse, error) {
	availCarriers, err := c.carrierRepository.GetAvailable()
	if err != nil {
		return nil, err
	}
	return utils.FilterPerformanceFromList(availCarriers), nil
}

type CarrierServiceInterface interface {
	GetAllCarriers() ([]simdomain.CarrierResponse, error)
	GetCarrierDetail(Id string) (*simdomain.CarrierResponse, error)
	GetAvailableCarriers() ([]simdomain.CarrierResponse, error)
}

func NewCarrierService(carrierRepo repository.CarrierRepositoryInterface) CarrierServiceInterface {
	return &CarrierService{
		carrierRepository: carrierRepo,
	}
}
