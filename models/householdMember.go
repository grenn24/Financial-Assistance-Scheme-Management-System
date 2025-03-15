package models

import (
	"time"

	"github.com/google/uuid"
)

type HouseholdMember struct {
	Base
	HouseholdOwnerID uuid.UUID     `gorm:"not null" json:"-"` //many-to-one
	Name             string        `gorm:"not null" json:"name" validate:"required"`
	EmploymentStatus *bool         `gorm:"not null" json:"employment_status" validate:"required"`
	MaritalStatus    MaritalStatus `gorm:"type:marital_status ; null" json:"marital_status" validate:"required,oneofci=single married widowed divorced"`
	Sex              Sex           `gorm:"type:sex ; not null" json:"sex" validate:"required,oneofci=male female"`
	DOB              time.Time     `gorm:"not null" json:"date_of_birth" validate:"required"`
	Relation         Relation      `gorm:"type:relation ; not null" json:"relation" validate:"required,oneofci=husband wife son daughter brother sister"`
	SchoolLevel      *SchoolLevel   `gorm:"type:school_level" json:"school_level,omitempty" validate:"omitempty,oneofci=primary secondary tertiary"`

	// Preloadable Columns
	HouseholdOwner *Applicant `gorm:"foreignKey:HouseholdOwnerID ; references:ID" json:"household_owner,omitempty"`
}

type Relation string

const (
	Husband  Relation = "husband"
	Wife     Relation = "wife"
	Son      Relation = "son"
	Daughter Relation = "daughter"
	Brother  Relation = "brother"
	Sister   Relation = "sister"
)

type SchoolLevel string

const (	
	Primary   SchoolLevel = "primary"
	Secondary SchoolLevel = "secondary"
	Tertiary  SchoolLevel = "tertiary"
)
