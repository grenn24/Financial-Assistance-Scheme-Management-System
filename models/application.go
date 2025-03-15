package models

import (
	"time"

	"github.com/google/uuid"
	_ "gorm.io/gorm"
)

type Application struct {
	Base
	ApplicantID uuid.UUID `gorm:"not null ; constraint:OnDelete:CASCADE" json:"applicant_id" validate:"required"` //many-to-one
	SchemeID    uuid.UUID `gorm:"not null ; constraint:OnDelete:CASCADE" json:"scheme_id" validate:"required"`    //many-to-one
	Status      Status    `gorm:"type:status ; not null" json:"status" validate:"required,oneof=pending approved rejected"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	// Preloadable Columns
	Applicant *Applicant `gorm:"foreignKey:ApplicantID ; references:ID ; constraint:OnDelete:CASCADE" json:"applicant"`
	Scheme    *Scheme    `gorm:"foreignKey:SchemeID ; references:ID ; constraint:OnDelete:CASCADE" json:"scheme"`
}

type Status string

const (
	Pending  Status = "pending"
	Approved Status = "approved"
	Rejected Status = "rejected"
)

type UpdateApplicationRequest struct {
	Base
	ApplicantID uuid.UUID `gorm:"not null ; constraint:OnDelete:CASCADE" json:"applicant_id" validate:"omitempty"` //many-to-one
	SchemeID    uuid.UUID `gorm:"not null ; constraint:OnDelete:CASCADE" json:"scheme_id" validate:"omitempty"`    //many-to-one
	Status      Status    `gorm:"type:status ; not null" json:"status" validate:"omitempty,oneof=pending approved rejected"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	// Preloadable Columns
	Applicant *Applicant `gorm:"foreignKey:ApplicantID ; references:ID ; constraint:OnDelete:CASCADE" json:"applicant"`
	Scheme    *Scheme    `gorm:"foreignKey:SchemeID ; references:ID ; constraint:OnDelete:CASCADE" json:"scheme"`
}

type CreateMultipleApplicationsRequest struct {
	Base
	ApplicantID uuid.UUID `gorm:"-" json:"applicant_id" validate:"required"`
	SchemeID    []uuid.UUID `gorm:"-" json:"scheme_id" validate:"required"`
	Status      Status    `gorm:"type:status ; not null" json:"status" validate:"required,oneof=pending approved rejected"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

