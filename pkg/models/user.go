package models

import (
	"time"
)

// User represents a platform user
type User struct {
	ID           string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Email        string    `gorm:"uniqueIndex;not null"`
	Password     string    `gorm:"not null"` // Hashed password
	APIKey       string    `gorm:"uniqueIndex"`
	Role         string    `gorm:"type:varchar(20);not null;default:'user'"`
	CreatedAt    time.Time `gorm:"not null;default:current_timestamp"`
	UpdatedAt    time.Time `gorm:"not null;default:current_timestamp"`
	LastLoginAt  *time.Time
	Organization string `gorm:"type:varchar(100)"`
}
