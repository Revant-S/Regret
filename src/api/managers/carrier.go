package managers

import (
	"Regret/domain"
	"Regret/src/api/storage"
)

type CarrierManager struct {
	carrierStorage storage.CarrierStorageInterface
}

func (c CarrierManager) UpdateCarriers(carriers []domain.Carrier) error {
	//TODO implement me
	panic("implement me")
}

func (c CarrierManager) UpdateCarrier(carrierId string, updatedCarrier *domain.Carrier) error {
	//TODO implement me
	panic("implement me")
}

type CarrierManagerInterface interface {
	UpdateCarriers(carriers []domain.Carrier) error
	UpdateCarrier(carrierId string, updatedCarrier *domain.Carrier) error
}

func NewCarrierManager(carrierStorage storage.CarrierStorageInterface) CarrierManagerInterface {
	return &CarrierManager{
		carrierStorage: carrierStorage,
	}
}
