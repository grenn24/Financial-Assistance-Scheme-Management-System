package models

import (
	"github.com/google/uuid"
	_ "gorm.io/gorm"
)

type SchemeCriteria struct {
	Base
	SchemeID uuid.UUID `gorm:"not null"` //one-to-one
	EmploymentStatus *bool `json:"employment_status,omitempty"`
	MaritalStatus *bool `json:"marital_status,omitempty"`
	HasChildren *HasChildren `gorm:"type:jsonb" json:"has_children,omitempty"`
	// Preloadable Columns
	Scheme *Scheme `gorm:"foreignKey:SchemeID ; references:ID"`
}

type HasChildren struct {
	SchoolLevel string `json:"school_level,omitempty"`
}