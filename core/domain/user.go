package domain

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey"`
	GoogleID  string    `gorm:"uniqueIndex;size:255;not null"`
	Name      string    `gorm:"size:255;not null"`
	Email     string    `gorm:"uniqueIndex;size:255;not null"`
	Avatar    string    `gorm:"size:512"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
