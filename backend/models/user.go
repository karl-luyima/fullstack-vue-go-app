package models

import "time"

// User is a database table, described as a Go struct. GORM reads the
// `gorm:"..."` tags to know how to build the actual SQL table, and the
// `json:"..."` tags control how this struct is converted to JSON when
// sent to the frontend.
type User struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Name         string    `gorm:"size:120;not null" json:"name"`
	Email        string    `gorm:"size:190;uniqueIndex;not null" json:"email"`
	PasswordHash string    `gorm:"size:255;not null" json:"-"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}