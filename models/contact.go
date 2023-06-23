package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Contact struct {
	Id        string `gorm:"primaryKey"`
	Name      string
	Email     string
	Phone     string
	Gender    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (contact *Contact) TableName() string {
	return "contacts"
}

func (contact *Contact) BeforeCreate(tx *gorm.DB) (err error) {
	contact.Id = uuid.New().String()

	return
}
