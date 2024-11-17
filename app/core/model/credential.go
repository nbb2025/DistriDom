package model

import (
	"github.com/google/uuid"
	"time"
)

type Credential struct {
	ID             uuid.UUID              `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserID         uuid.UUID              `gorm:"type:uuid;not null"`
	Provider       string                 `gorm:"not null"`
	CredentialData map[string]interface{} `gorm:"type:jsonb;not null"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
