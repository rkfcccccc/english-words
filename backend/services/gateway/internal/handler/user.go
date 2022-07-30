package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	user_service "github.com/rkfcccccc/english_words/services/gateway/internal/service/user"
	"github.com/rkfcccccc/english_words/services/gateway/internal/service/verification"
	"github.com/rkfcccccc/english_words/services/gateway/pkg/auth"
)

type signupInput struct {
	Email        string             `json:"email" binding:"required"`
	Password     string             `json:"password" binding:"required"`
	Verification *verificationInput `json:"verification"`
}

func (h *Handlers) issueCredentials(c *gin.Context, userId int) (string, string, error) {
	jwt, err := h.Auth.IssueJWT(userId)
	if err != nil {
		return "", "", fmt.Errorf("Auth.IssueJWT: %v", err)
	}

	refresh, err := h.Auth.IssueRefreshToken(c, userId)
	if err != nil {
		return "", "", fmt.Errorf("Auth.IssueRefreshToken: %v", err)
	}

	return jwt, refresh, nil
}

func (h *Handlers) UserSignup(c *gin.Context) {
	var body signupInput

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	proceed, err := h.verifyRequest(c, verification.Registration, body.Email, body.Verification)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	} else if !proceed {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	userId, err := h.Services.User.Create(c, body.Email, body.Password)

	switch err {
	case user_service.ErrAlreadyExists:
		c.AbortWithStatus(http.StatusBadRequest)
		return
	case user_service.ErrInvalidEmail:
		c.AbortWithStatus(http.StatusBadRequest)
		return
	case user_service.ErrTooLongPassword:
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("User.Create: %v", err))
		return
	}

	jwt, err := h.Auth.IssueJWT(userId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Auth.IssueJWT: %v", err))
		return
	}

	refresh, err := h.Auth.IssueRefreshToken(c, userId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Auth.IssueRefreshToken: %v", err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"jwt":     jwt,
		"refresh": refresh,
	})
}

type refreshInput struct {
	Token string `json:"token" binding:"required"`
}

func (h *Handlers) UserRefresh(c *gin.Context) {
	var body refreshInput

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userId, err := h.Auth.GetRefreshToken(c, body.Token)
	if errors.Is(err, auth.ErrRefreshMiss) {
		c.AbortWithStatus(http.StatusNotFound)
		return
	} else if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Auth.GetRefreshToken: %v", err))
		return
	}

	jwt, refresh, err := h.issueCredentials(c, userId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if err := h.Auth.DeleteRefreshToken(c, body.Token); err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Auth.DeleteRefreshToken: %v", err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"jwt":     jwt,
		"refresh": refresh,
	})
}

type loginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// TODO: maybe somehow restrict max sessions per user
func (h *Handlers) UserLogin(c *gin.Context) {
	var body loginInput

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user, err := h.Services.User.GetByEmailAndPassword(c, body.Email, body.Password)
	if errors.Is(err, user_service.ErrNotFound) {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	} else if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	jwt, refresh, err := h.issueCredentials(c, user.Id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"jwt":     jwt,
		"refresh": refresh,
	})
}

type recoveryInput struct {
	Email        string             `json:"email" binding:"required"`
	Password     string             `json:"password" binding:"required"`
	Verification *verificationInput `json:"verification"`
}

func (h *Handlers) UserRecovery(c *gin.Context) {
	var body recoveryInput

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user, err := h.Services.User.GetByEmail(c, body.Email)
	if errors.Is(err, user_service.ErrNotFound) {
		c.AbortWithStatus(http.StatusNotFound)
		return
	} else if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	proceed, err := h.verifyRequest(c, verification.Registration, body.Email, body.Verification)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	} else if !proceed {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// TODO: h.Services.User.UpdatePassword(c, user.Id, body.Password) function
	panic("UserRecovery unimlemented")

	jwt, refresh, err := h.issueCredentials(c, user.Id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"jwt":     jwt,
		"refresh": refresh,
	})
}
