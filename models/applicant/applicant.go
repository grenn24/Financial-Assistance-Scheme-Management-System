package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Applicant struct {
	gorm.Model
	ID               uuid.UUID   `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name             string      `json:"name"`
	EmploymentStatus bool        `json:"employment_status"`
	MaritalStatus    string      `gorm:"type:ENUM('Single', 'Married', 'Widowed', 'Divorced')" json:"marital_status"`
	Sex              string      `gorm:"type:ENUM('Male', 'Female')" json:"sex"`
	DOB              time.Time   `json:"date_of_birth"`
	ParentID         *uuid.UUID  //many-to-one
	SpouseID         *uuid.UUID  //one-to-one
	Household        []Applicant `json:"household"`
	Relation         *string     `gorm:"type:ENUM('Husband', 'Wife', 'Son', 'Daughter')"`
}
