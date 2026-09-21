package managers

import (
	"Regret/domain"
	"Regret/src/api/storage"
)

type DispatchManager struct {
	carrierStore storage.CarrierStorageInterface
}

type DispatchManagerInterface interface {
	GetCarrier(message *domain.Message) (*storage.CarrierStorage, error)
}

func NewDispatchManager(carrierStorage storage.CarrierStorageInterface) DispatchManagerInterface {
	return &DispatchManager{
		carrierStore: carrierStorage,
	}
}

func (dm *DispatchManager) GetCarrier(message *domain.Message) (*storage.CarrierStorage, error) {
	return nil, nil
}
