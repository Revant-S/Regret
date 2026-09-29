package managers

import (
	"Regret/domain"
	"Regret/src/api/storage"
)

type carrierManager struct {
	carrierStorage storage.CarrierStorageInterface
}

func (c carrierManager) UpdateCarriers(carriers []domain.Carrier) error {
	//TODO implement me
	panic("implement me")
}

func (c carrierManager) UpdateCarrier(carrierId string, updatedCarrier *domain.Carrier) error {
	//TODO implement me
	panic("implement me")
}

type CarrierManager interface {
	UpdateCarriers(carriers []domain.Carrier) error
	UpdateCarrier(carrierId string, updatedCarrier *domain.Carrier) error
}

func NewCarrierManager(carrierStorage storage.CarrierStorageInterface) CarrierManager {
	return &carrierManager{
		carrierStorage: carrierStorage,
	}
}
