package models

import (
	"time"

	_ "github.com/google/uuid"
	_ "gorm.io/gorm"
)

type Applicant struct {
	Base
	Name             string        `gorm:"not null" json:"name" validate:"required"`
	EmploymentStatus *bool         `gorm:"not null" json:"employment_status" validate:"required"`
	MaritalStatus    MaritalStatus `gorm:"type:marital_status ; null" json:"marital_status" validate:"required,oneof=single married widowed divorced"`
	Sex              Sex           `gorm:"type:sex ; not null" json:"sex" validate:"required,oneof=male female"`
	DOB              time.Time     `gorm:"not null" json:"date_of_birth" validate:"required"`

	// Preloadable Columns
	Household []HouseholdMember `gorm:"foreignKey:HouseholdOwnerID ; references:ID ; constraint:onDelete:CASCADE" json:"household" validate:"omitempty,dive"`
}

type MaritalStatus string

const (
	Single   MaritalStatus = "single"
	Married  MaritalStatus = "married"
	Widowed  MaritalStatus = "widowed"
	Divorced MaritalStatus = "divorced"
)



type Sex string

const (
	Male   Sex = "male"
	Female Sex = "female"
)

type UpdateApplicantRequest struct {
	Base
	Name             string        `gorm:"not null" json:"name" validate:"omitempty"`
	EmploymentStatus *bool         `gorm:"not null" json:"employment_status" validate:"omitempty"`
	MaritalStatus    MaritalStatus `gorm:"type:marital_status ; null" json:"marital_status" validate:"omitempty,oneof=single married widowed divorced"`
	Sex              Sex           `gorm:"type:sex ; not null" json:"sex" validate:"omitempty,oneof=male female"`
	DOB              time.Time     `gorm:"not null" json:"date_of_birth" validate:"omitempty"`

	// Preloadable Columns
	Household []HouseholdMember `gorm:"foreignKey:HouseholdOwnerID ; references:ID ; constraint:onDelete:CASCADE" json:"household" validate:"omitempty,dive"`
}