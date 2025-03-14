package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Base model containing common columns
type Base struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BaseHidden struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primary_key" json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (base *Base) BeforeCreate(db *gorm.DB) error {
	uuid, err := uuid.NewV7()
	if err != nil {
		return err
	}
	base.ID = uuid
	return nil
}


func (base *BaseHidden) BeforeCreate(db *gorm.DB) error {
	uuid, err := uuid.NewV7()
	if err != nil {
		return err
	}
	base.ID = uuid
	return nil
}

