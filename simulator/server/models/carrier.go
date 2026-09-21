package models

import (
	"Regret/domain"
	"gorm.io/gorm"
)

type Carrier struct {
	gorm.Model
	domain.Carrier

	Reliability float64 `gorm:"reliability"`
	FailureRate float64 `gorm:"failure_rate"`
}
