package models

import "time"

type User struct {
	ID           uint      `gorm:"primaryKey"`
	Email        string    `gorm:"unique"`
	Password     string
	Role         string `gorm:"default:user"` // user | admin
	ResetToken   *string
	ResetExpires *time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}