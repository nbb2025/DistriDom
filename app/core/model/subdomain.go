package model

import (
	"github.com/google/uuid"
	"time"
)

type Subdomain struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	DomainID      uuid.UUID `gorm:"type:uuid;not null"`
	SubdomainName string    `gorm:"unique;not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
