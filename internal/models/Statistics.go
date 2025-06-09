package models

import "time"

type Statistics struct {
	StatisticsId int
	UserId       int

	// TODO add statistics related fields

	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
