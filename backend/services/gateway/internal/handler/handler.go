package handler

import (
	"github.com/rkfcccccc/english_words/services/gateway/internal/service"
	"github.com/rkfcccccc/english_words/services/gateway/pkg/auth"
)

type Handlers struct {
	Services *service.Services
	Auth     *auth.Helper
}

func NewHandlers(services *service.Services, authHelper *auth.Helper) *Handlers {
	return &Handlers{services, authHelper}
}

func (h *Handlers) newError(errName string) map[string]any {
	return map[string]any{"error_name": errName}
}
