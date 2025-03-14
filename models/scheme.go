package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/google/uuid"
	_ "github.com/google/uuid"
	_ "gorm.io/gorm"
)

type Scheme struct {
	Base
	Name        string          `gorm:"not null" json:"name" validate:"required"`
	Description string          `gorm:"not null" json:"description"  validate:"required"`
	// Preloadable Columns (on delete cascade)
	Benefits    []SchemeBenefit `gorm:"foreignKey:SchemeID ; references:ID ; constraint:OnDelete:CASCADE" json:"benefits" validate:"required,dive"` //one-to-many
	Criteria    SchemeCriteria  `gorm:"foreignKey:SchemeID ; references:ID ; constraint:OnDelete:CASCADE" json:"criteria" validate:"required"` //one-to-one
}

type SchemeBenefit struct {
	Base
	SchemeID uuid.UUID `gorm:"not null"` //many-to-one
	Name     string    `gorm:"not null" json:"name" validate:"required"`
	Amount   float64   `json:"amount,omitempty"`
	// Preloadable Columns
	Scheme *Scheme `gorm:"foreignKey:SchemeID ; references:ID" json:"scheme,omitempty"`
}

type SchemeCriteria struct {
	Base
	SchemeID         uuid.UUID      `gorm:"not null"` //one-to-one
	EmploymentStatus *bool          `json:"employment_status,omitempty"`
	MaritalStatus    *MaritalStatus `json:"marital_status,omitempty" validate:"oneof=single married widowed divorced"`
	HasChildren      *HasChildren   `gorm:"type:jsonb" json:"has_children,omitempty"`

	// Preloadable Columns
	Scheme *Scheme `gorm:"foreignKey:SchemeID ; references:ID ; onDelete:CASCADE" json:"scheme,omitempty"`
}

type HasChildren struct {
	Status       bool `json:"status" validate:"required"`
	SchoolLevel string `json:"school_level,omitempty" validate:"oneof=primary secondary tertiary all"`
}

// Unmarshal JSONB (bytes) into HasChildren struct
func (hasChildren *HasChildren) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	// Unmarshal the bytes directly into the HasChildren struct
	return json.Unmarshal(bytes, &hasChildren)
}

// Marshal HasChildren struct into JSONB (bytes)
func (hasChildren HasChildren) Value() (driver.Value, error) {
	if !hasChildren.Status {
		return nil, nil
	}
	return json.Marshal(hasChildren)
}
