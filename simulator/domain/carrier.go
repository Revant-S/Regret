package domain

type CarrierStatus string

const (
	CarrierAvailable   CarrierStatus = "CarrierAvailable"
	CarrierUnAvailable CarrierStatus = "CarrierUnAvailable"
	CarrierInvalid     CarrierStatus = "CarrierInvalid"
)

type Carrier struct {
	Id          string
	Cost        float64
	Reliability float64
	Status      CarrierStatus
}

type CarrierResponse struct {
	Id     string        `json:"Id"`
	Cost   float64       `json:"Cost"`
	Status CarrierStatus `json:"Status"`
}

type CarrierInterface interface {
	Send(message Message) error
}

func NewCarrier(cost float64, reliability float64) CarrierInterface {
	return &Carrier{Cost: cost, Reliability: reliability}
}

func (c *Carrier) Send(message Message) error {
	return nil
}
