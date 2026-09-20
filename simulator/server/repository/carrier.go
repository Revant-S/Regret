package repository

import (
	"Regret/simulator/domain"
	"errors"
	"slices"
)

type CarrierRepository struct {
}

func (cr *CarrierRepository) GetAll() ([]domain.Carrier, error) {
	return DummyCarriers, nil
}

func (cr *CarrierRepository) GetAvailable() ([]domain.Carrier, error) {
	var active []domain.Carrier
	for _, carrier := range DummyCarriers {
		if carrier.Status == domain.CarrierAvailable {
			active = append(active, carrier)
		}
	}
	return active, nil
}

func (cr *CarrierRepository) GetDetails(Id string) (*domain.Carrier, error) {

	carrierIdx := slices.IndexFunc(DummyCarriers, func(carrier domain.Carrier) bool {
		return carrier.Id == Id
	})
	if carrierIdx == -1 {
		return nil, errors.New("carrier not found")
	}
	return &DummyCarriers[carrierIdx], nil
}

type CarrierRepositoryInterface interface {
	GetAll() ([]domain.Carrier, error)
	GetAvailable() ([]domain.Carrier, error)
	GetDetails(Id string) (*domain.Carrier, error)
}

func NewCarrierRepository() CarrierRepositoryInterface {
	return &CarrierRepository{}
}
