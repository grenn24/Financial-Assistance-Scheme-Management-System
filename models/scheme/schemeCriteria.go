package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SchemeCriteria struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	SchemeID uuid.UUID 
	EmploymentStatus bool
	MaritalStatus bool
}