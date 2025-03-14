package models

import (
	"time"

	"github.com/google/uuid"
)

type HouseholdMember struct {
	Base
	HouseholdOwnerID uuid.UUID     `gorm:"not null" json:"household_owner_id"` //many-to-one
	Name             string        `gorm:"not null" json:"name" validate:"required"`
	EmploymentStatus *bool         `gorm:"not null" json:"employment_status" validate:"required"`
	MaritalStatus    MaritalStatus `gorm:"type:marital_status ; null" json:"marital_status" validate:"oneofci=single married widowed divorced"`
	Sex              Sex           `gorm:"type:sex ; not null" json:"sex" validate:"oneofci=male female"`
	DOB              time.Time     `gorm:"not null" json:"date_of_birth" validate:"required"`
	Relation         Relation      `gorm:"type:relation ; not null" json:"relation" validate:"oneofci=husband wife son daughter"`

	// Preloadable Columns
	HouseholdOwner *Applicant `gorm:"foreignKey:HouseholdOwnerID ; references:ID" json:"household_owner,omitempty"`
}

type Relation string

const (
	Husband  Relation = "husband"
	Wife     Relation = "wife"
	Son      Relation = "son"
	Daughter Relation = "daughter"
)
