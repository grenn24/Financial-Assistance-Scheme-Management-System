package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SchemeBenefit struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	SchemeID uuid.UUID
	Name     string
	Amount   float64
}