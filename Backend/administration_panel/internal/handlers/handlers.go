package handlers

import (
	"github.com/Artmoond/Minion-Team-TTK-Case/Backend/administration_panel/internal/service"
)

type handlers struct {
	service service.PanelService
}

func NewHandlers(service service.PanelService) *handlers {
	return &handlers{service: service}
}
