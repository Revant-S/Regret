package policy

import (
	"Regret/domain"
	"Regret/src/api/storage"
)

type BasicPolicy struct {
	storage storage.CarrierStorageInterface
}

func (b *BasicPolicy) GetBestChannel(message *domain.Message) domain.Carrier {
	//TODO implement me
	panic("implement me")
}

type BasicPolicyInterface interface {
	GetBestChannel(message *domain.Message) domain.Carrier
}

func NewBasicPolicy(storageInterface storage.CarrierStorageInterface) BasicPolicyInterface {
	return &BasicPolicy{
		storage: storageInterface,
	}
}
