package database

import "gorm.io/gorm"

type Database struct {
}

func NewDatabase(conn *gorm.DB) *Database {
	return &Database{}
}
