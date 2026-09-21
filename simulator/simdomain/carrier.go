package simdomain

import "Regret/domain"

type Carrier struct {
	domain.Carrier

	Reliability float64
	FailureRate float64
}

type CarrierResponse struct {
	domain.Carrier
}

type CarrierInterface interface {
	Send(message domain.Message) error
}

func NewCarrier(cost float64, reliability float64) CarrierInterface {
	return &Carrier{
		Carrier: domain.Carrier{
			Cost: cost,
		},
		Reliability: reliability,
	}
}

func (c *Carrier) Send(message domain.Message) error {
	return nil
}
