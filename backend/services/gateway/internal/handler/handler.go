package handler

import (
	"github.com/rkfcccccc/english_words/services/gateway/pkg/auth"
	"github.com/rkfcccccc/english_words/shared_pkg/services"
)

type Handlers struct {
	Services *services.Services
	Auth     *auth.Helper
}

func NewHandlers(services *services.Services, authHelper *auth.Helper) *Handlers {
	return &Handlers{services, authHelper}
}

func (h *Handlers) newError(errName string) map[string]any {
	return map[string]any{"error_name": errName}
}
