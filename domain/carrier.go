package domain

type CarrierStatus string

const (
	CarrierAvailable   CarrierStatus = "CarrierAvailable"
	CarrierUnAvailable CarrierStatus = "CarrierUnAvailable"
	CarrierInvalid     CarrierStatus = "CarrierInvalid"
)

type Carrier struct {
	Id     string        `json:"Id"`
	Cost   float64       `json:"Cost"`
	Status CarrierStatus `json:"Status"`
}
