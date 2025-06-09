package models

import "time"

// TardinessReason - причина опоздания
type TardinessReason string

// Tardiness - Опоздание.
type Tardiness struct {
	TardinessId int

	ExpectedArrivalTime time.Time
	Reason              TardinessReason
	Comment             string

	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

// DayOff - отгул.
type DayOff struct {
	DayOffId int

	From    time.Time
	To      time.Time
	Reason  TardinessReason
	Comment string

	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

// Vacation - отпуск.
type Vacation struct {
	VacationId int

	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

// Sickness - болезнь.
type Sickness struct {
	SicknessId int

	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

// BusinessTrip - командировка
type BusinessTrip struct {
	BusinessTripId int

	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

// WorkRelatedLeave - уход по работе/рабочим задачам.
type WorkRelatedLeave struct {
	WorkRelatedLeaveId int

	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

// Overtime - переработки.
type Overtime struct {
	OvertimeId int

	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

// OtherEvent - другое.
// Для событий не подходящим ни к одному типу.
type OtherEvent struct {
	OtherEventId int

	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
