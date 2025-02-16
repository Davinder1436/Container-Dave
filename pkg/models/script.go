package models

import "time"

// Script represents an AI-generated script
type Script struct {
	ID          string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID      string    `gorm:"type:uuid;not null"`
	ContainerID string    `gorm:"type:uuid"`
	Name        string    `gorm:"not null"`
	Content     string    `gorm:"type:text;not null"`
	Language    string    `gorm:"type:varchar(50);not null"`
	Version     int       `gorm:"not null;default:1"`
	Status      string    `gorm:"type:varchar(20);not null;default:'pending'"`
	CreatedAt   time.Time `gorm:"not null;default:current_timestamp"`
	UpdatedAt   time.Time `gorm:"not null;default:current_timestamp"`
	LastRunAt   *time.Time
	Metadata    JSONMap `gorm:"type:jsonb"`
}

// ScriptExecution represents a single execution of a script
type ScriptExecution struct {
	ID          string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	ScriptID    string    `gorm:"type:uuid;not null"`
	ContainerID string    `gorm:"type:uuid;not null"`
	StartTime   time.Time `gorm:"not null;default:current_timestamp"`
	EndTime     *time.Time
	Status      string `gorm:"type:varchar(20);not null"`
	ExitCode    int
	Output      string  `gorm:"type:text"`
	Error       string  `gorm:"type:text"`
	Resources   JSONMap `gorm:"type:jsonb"` // CPU, memory usage etc.
}
