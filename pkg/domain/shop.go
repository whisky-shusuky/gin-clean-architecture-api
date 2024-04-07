package domain

import "time"

type Shop struct {
	ID              uint64     `gorm:"primary_key"`
	ShopName        string     `gorm:"not null"`
	ShopDescription string     `gorm:"not null"`
	CreatedAt       time.Time  `gorm:"not null"`
	DeletedAt       *time.Time `gorm:"index"`
}
