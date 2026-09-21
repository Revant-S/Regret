package storage

import "Regret/domain"

type CarrierStorage struct {
}

type CarrierStorageInterface interface {
	GetAllCarrier() []domain.Carrier
}

func NewCarrierStorage() CarrierStorageInterface {
	return &CarrierStorage{}
}

func (c *CarrierStorage) GetAllCarrier() []domain.Carrier {
	panic("Not Implemented GetAll Carriers")
}
