package schema

import (
	"Regret/domain"
	"gorm.io/gorm"
)

type Carrier struct {
	gorm.Model
	domain.Carrier
}
