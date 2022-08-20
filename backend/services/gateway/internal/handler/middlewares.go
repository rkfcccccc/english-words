package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rkfcccccc/english_words/services/gateway/pkg/auth"
)

func (h *Handlers) AuthRequired(c *gin.Context) {
	header := c.GetHeader("Authorization")

	parts := strings.Split(header, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, h.newError("INVALID_TOKEN"))
		return
	}

	claims, err := h.Auth.ParseJWT(parts[1])
	switch err {
	case auth.ErrInvalidToken:
		c.AbortWithStatusJSON(http.StatusUnauthorized, h.newError("INVALID_TOKEN"))
		return
	case auth.ErrTokenExpired:
		c.AbortWithStatusJSON(http.StatusUnauthorized, h.newError("TOKEN_EXPIRED"))
		return
	}

	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	c.Set("user_id", claims.UserId)
}

func (h *Handlers) ServiceKeyRequired(c *gin.Context) {
	header := c.GetHeader("Service-Key")

	if !h.Auth.CheckServiceKey(header) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, h.newError("INVALID_KEY"))
		return
	}
}
