package models

import (
	"time"

	"github.com/google/uuid"
	_ "gorm.io/gorm"
)

type Application struct {
	Base
	ApplicantID uuid.UUID `gorm:"not null" json:"applicant_id" validate:"required"` //many-to-one
	SchemeID    uuid.UUID `gorm:"not null" json:"scheme_id" validate:"required"` //many-to-one
	Status      Status    `gorm:"type:status ; not null" json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	// Preloadable Columns
	Applicant *Applicant `gorm:"foreignKey:ApplicantID ; references:ID"`
	Scheme    *Scheme    `gorm:"foreignKey:SchemeID ; references:ID"`
}

type Status string

const (
	Pending  Status = "pending"
	Approved Status = "approved"
	Rejected Status = "rejected"
)
