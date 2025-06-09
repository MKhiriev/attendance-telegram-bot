package handlers

import (
	"attendance-telegram-bot/internal/services"
)

type Handlers struct {
	service *services.Service
}

func NewHandler(s *services.Service) *Handlers {
	return &Handlers{service: s}
}

// Init add handlers to specific routes
func (h *Handlers) Init() {
	// your code here
}
