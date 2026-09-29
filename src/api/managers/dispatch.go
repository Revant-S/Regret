package managers

import (
	"Regret/domain"
	"Regret/src/api/storage"
)

type dispatchManager struct {
	carrierStore storage.CarrierStorageInterface
}

type DispatchManager interface {
	GetCarrier(message *domain.Message) (*storage.CarrierStorage, error)
}

func NewDispatchManager(carrierStorage storage.CarrierStorageInterface) DispatchManager {
	return &dispatchManager{
		carrierStore: carrierStorage,
	}
}

func (dm *dispatchManager) GetCarrier(message *domain.Message) (*storage.CarrierStorage, error) {
	return nil, nil
}
