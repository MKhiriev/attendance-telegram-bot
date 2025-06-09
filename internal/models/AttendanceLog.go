package models

import "time"

type AttendanceLog struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	Date      time.Time
	EnterTime time.Time
	ExitTime  time.Time // Очень сомневаюсь, что СКУД способен выход считать, но хай пусть будет
}
