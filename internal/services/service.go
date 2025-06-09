package services

import (
	"attendance-telegram-bot/internal/database"
	"attendance-telegram-bot/internal/models"
)

type Service struct {
	Database *database.Database
}

func NewService(db *database.Database) *Service {
	return &Service{
		Database: db,
	}
}

// TardinessService сервис по регистрации опозданий
type TardinessService interface {
	CreateTardiness(tardinessEvent models.Tardiness) (models.Tardiness, error)
	GetTardinessById(tardinessEventId int) (models.Tardiness, error)
	GetAllTardinessEvents() ([]models.Tardiness, error)
	UpdateTardiness(tardinessEvent models.Tardiness) (models.Tardiness, error)
	DeleteTardiness(tardinessEventId int) (bool, error)
}

// TODO написать интерфейсы сервисов
