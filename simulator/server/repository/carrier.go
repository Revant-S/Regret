package repository

import (
	"Regret/domain"
	"Regret/simulator/simdomain"
	"errors"
	"gorm.io/gorm"
	"slices"
)

type CarrierRepository struct {
	DB *gorm.DB
}

func (cr *CarrierRepository) GetAll() ([]simdomain.Carrier, error) {
	return DummyCarriers, nil
}

func (cr *CarrierRepository) GetAvailable() ([]simdomain.Carrier, error) {
	var active []simdomain.Carrier
	for _, carrier := range DummyCarriers {
		if carrier.Status == domain.CarrierAvailable {
			active = append(active, carrier)
		}
	}
	return active, nil
}

func (cr *CarrierRepository) GetDetails(Id string) (*simdomain.Carrier, error) {

	carrierIdx := slices.IndexFunc(DummyCarriers, func(carrier simdomain.Carrier) bool {
		return carrier.Name == Id
	})
	if carrierIdx == -1 {
		return nil, errors.New("carrier not found")
	}
	return &DummyCarriers[carrierIdx], nil
}

type CarrierRepositoryInterface interface {
	GetAll() ([]simdomain.Carrier, error)
	GetAvailable() ([]simdomain.Carrier, error)
	GetDetails(Id string) (*simdomain.Carrier, error)
}

func NewCarrierRepository(DB *gorm.DB) CarrierRepositoryInterface {
	return &CarrierRepository{
		DB: DB,
	}
}
