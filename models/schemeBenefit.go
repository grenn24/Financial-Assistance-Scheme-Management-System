package models

import (
	"github.com/google/uuid"
	_ "gorm.io/gorm"
)

type SchemeBenefit struct {
	Base
	SchemeID uuid.UUID `gorm:"not null"`
	Name     string    `gorm:"not null" json:"name" validate:"required"`
	Amount   float64   `json:"amount,omitempty"`
}
