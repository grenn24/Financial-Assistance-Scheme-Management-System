package models

import (
	"time"

	"github.com/google/uuid"
	_ "gorm.io/gorm"
)

type Applicant struct {
	Base
	Name             string        `gorm:"not null" json:"name" validate:"required"`
	EmploymentStatus *bool         `gorm:"not null" json:"employment_status" validate:"required"`
	MaritalStatus    MaritalStatus `gorm:"type:marital_status ; null" json:"marital_status" validate:"oneofci=single married widowed divorced"`
	Sex              Sex           `gorm:"type:sex ; not null" json:"sex" validate:"oneofci=male female"`
	DOB              time.Time     `gorm:"not null" json:"date_of_birth" validate:"required"`
	ParentID         *uuid.UUID    //many-to-one
	SpouseID         *uuid.UUID    //one-to-one
	Relation         *Relation     `gorm:"type:relation" json:"relation,omitempty"`
}

type MaritalStatus string

const (
	Single   MaritalStatus = "single"
	Married  MaritalStatus = "married"
	Widowed  MaritalStatus = "widowed"
	Divorced MaritalStatus = "divorced"
)

type Relation string

const (
	Husband  Relation = "husband"
	Wife     Relation = "wife"
	Son      Relation = "son"
	Daughter Relation = "daughter"
)

type Sex string

const (
	Male   Sex = "male"
	Female Sex = "female"
)
