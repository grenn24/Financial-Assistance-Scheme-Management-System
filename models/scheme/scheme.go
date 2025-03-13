package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Scheme struct {
	gorm.Model
	ID          uuid.UUID       `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Benefits    []SchemeBenefit `gorm:"foreignKey:SchemeID" json:"benefits"` //one-to-many
	Criteria    SchemeCriteria  `gorm:"foreignKey:SchemeID" json:"criteria"` //one-to-one
}
