package impl

type Carrier struct {
	Cost        int16
	Reliability float64
}

type CarrierInterface interface {
	Send(message Message) error
}

func NewCarrier(cost int16, reliability float64) CarrierInterface {
	return &Carrier{Cost: cost, Reliability: reliability}
}

func (c *Carrier) Send(message Message) error {
	return nil
}
