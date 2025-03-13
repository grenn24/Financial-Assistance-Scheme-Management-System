package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Application struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	ApplicantID uuid.UUID 
	SchemeID    uuid.UUID 
	Status      string    `gorm:"type:ENUM('Pending', 'Approved', 'Rejected')" json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
