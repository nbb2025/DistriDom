package model

import (
	"github.com/google/uuid"
	"time"
)

type Domain struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserID     uuid.UUID `gorm:"type:uuid;not null"`
	DomainName string    `gorm:"unique;not null"`
	IsTopLevel bool      `gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
