package models

import "time"

type Team struct {
	TeamId   int
	LeaderId int // check gorm.io if I can use structure as a field
	Name     string

	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
