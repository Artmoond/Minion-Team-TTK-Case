package handlers

import "github.com/Artmoond/Minion-Team-TTK-Case/internal/service"

type Handlers struct {
	s service.Service
}

func NewHandlers(s service.Service) *Handlers {
	return &Handlers{
		s: s,
	}
}
