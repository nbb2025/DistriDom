package model

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Username  string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	Email     string    `gorm:"unique;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
