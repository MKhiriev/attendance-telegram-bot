package models

import (
	"time"
)

type User struct {
	ID         uint   `gorm:"primaryKey"`
	TelegramID int64  `gorm:"uniqueIndex"`
	Username   string // username telegram
	FullName   string // ФИО сотрудника

	TeamId int // Команда разработки

	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
