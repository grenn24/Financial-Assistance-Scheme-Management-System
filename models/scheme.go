package models

import (
	_ "github.com/google/uuid"
	_ "gorm.io/gorm"
)

type Scheme struct {
	Base
	Name        string          `gorm:"not null" json:"name" validate:"required"`
	Description string          `gorm:"not null" json:"description"  validate:"required"`
	// Preloadable Columns
	Benefits    []SchemeBenefit `gorm:"foreignKey:SchemeID ; references:ID ; constraint:OnDelete:CASCADE" json:"benefits"` //one-to-many
	Criteria    SchemeCriteria  `gorm:"foreignKey:SchemeID ; references:ID ; constraint:OnDelete:CASCADE" json:"criteria"` //one-to-one
}
