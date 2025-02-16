package models

import (
	"time"

	"github.com/lib/pq"
)

// APIKey represents API keys for authentication
type APIKey struct {
	ID        string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID    string `gorm:"type:uuid;not null"`
	Key       string `gorm:"uniqueIndex;not null"`
	Name      string `gorm:"not null"`
	LastUsed  *time.Time
	ExpiresAt *time.Time
	Scopes    pq.StringArray `gorm:"type:text[]"`
	CreatedAt time.Time      `gorm:"not null;default:current_timestamp"`
}

// Audit represents audit trail entries
type Audit struct {
	ID         string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID     string    `gorm:"type:uuid"`
	Action     string    `gorm:"type:varchar(50);not null"`
	Resource   string    `gorm:"type:varchar(50);not null"`
	ResourceID string    `gorm:"type:uuid"`
	Timestamp  time.Time `gorm:"not null;default:current_timestamp"`
	Changes    JSONMap   `gorm:"type:jsonb"`
	IPAddress  string
}

// Metric represents system metrics
type Metric struct {
	ID        string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name      string    `gorm:"not null"`
	Value     float64   `gorm:"not null"`
	Unit      string    `gorm:"type:varchar(20)"`
	Labels    JSONMap   `gorm:"type:jsonb"`
	Timestamp time.Time `gorm:"not null;default:current_timestamp"`
}

// JSONMap is a helper type for JSON fields
type JSONMap map[string]interface{}
