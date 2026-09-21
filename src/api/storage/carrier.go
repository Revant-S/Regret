package storage

import (
	"Regret/domain"
	"gorm.io/gorm"
)

type CarrierStorage struct {
	DB *gorm.DB
}

type CarrierStorageInterface interface {
	GetAllCarrier() []domain.Carrier
}

func NewCarrierStorage(DB *gorm.DB) CarrierStorageInterface {
	return &CarrierStorage{
		DB: DB,
	}
}

func (c *CarrierStorage) GetAllCarrier() []domain.Carrier {
	panic("Not Implemented GetAll Carriers")
}
