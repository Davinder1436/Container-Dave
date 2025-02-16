package models

import (
	"time"

	"github.com/lib/pq"
)

// Container represents a Docker container instance
type Container struct {
	ID             string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID         string `gorm:"type:uuid;not null"`
	Name           string `gorm:"not null"`
	DockerID       string `gorm:"uniqueIndex"`
	Image          string `gorm:"not null"`
	Status         string `gorm:"type:varchar(20);not null"`
	IPAddress      string
	Port           int
	CreatedAt      time.Time `gorm:"not null;default:current_timestamp"`
	StartedAt      *time.Time
	TerminatedAt   *time.Time
	Environment    pq.StringArray `gorm:"type:text[]"`
	ResourceLimits JSONMap        `gorm:"type:jsonb"`
	HealthStatus   string         `gorm:"type:varchar(20);default:'unknown'"`
}

// Log represents system logs
type Log struct {
	ID        string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Level     string    `gorm:"type:varchar(10);not null"`
	Source    string    `gorm:"type:varchar(50);not null"`
	Message   string    `gorm:"type:text;not null"`
	Metadata  JSONMap   `gorm:"type:jsonb"`
	Timestamp time.Time `gorm:"not null;default:current_timestamp"`
}
