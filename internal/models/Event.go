package models

import "time"

// Event - запись о событии.
//
// Включает в себя:
//   - [Tardiness] - опоздание
//   - [DayOff] - отгул
//   - [Vacation] - отпуск
//   - [Sickness] - болезнь
//   - [BusinessTrip] - командировка
//   - [WorkRelatedLeave] - уход по работе/рабочим задачам
//   - [Overtime] - переработки
//   - [OtherEvent] - другое
type Event struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint //f.k на User
	EventType uint // late, sick, vacation, etc. f.k на EventTypes
	Comment   string
	Minutes   *int // only for late
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	DateFrom  time.Time //Начало периода. Дата события не обязательно совпадает с датой создания, например чел заранее предупредил за 1-2 дня
	DateTo    time.Time //Конец периода
	User      User      `gorm:"foreignKey:UserID"`
}
