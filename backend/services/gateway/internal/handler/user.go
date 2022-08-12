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

func (h *Handlers) canCreateAccount(c *gin.Context, email, password string) bool {
	ok, err := h.Services.User.CanCreate(c, email, password)

	switch err {
	case user_service.ErrAlreadyExists:
		c.AbortWithStatusJSON(http.StatusConflict, h.newError("ALREADY_EXISTS"))
		return false
	case user_service.ErrInvalidEmail:
		c.AbortWithStatusJSON(http.StatusBadRequest, h.newError("INVALID_EMAIL"))
		return false
	case user_service.ErrInvalidPassword:
		c.AbortWithStatusJSON(http.StatusBadRequest, h.newError("INVALID_PASSWORD"))
		return false
	}

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("User.CanCreate: %v", err))
		return false
	} else if !ok {
		c.AbortWithError(http.StatusInternalServerError, errors.New("User.CanCreate: cannot create an account"))
		return false
	}

	return true
}

func (h *Handlers) UserSignup(c *gin.Context) {
	var body signupInput

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if ok := h.canCreateAccount(c, body.Email, body.Password); !ok {
		return
	}

	proceed, err := h.verifyRequest(c, verification.Registration, body.Email, body.Verification)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	} else if !proceed {
		return
	}

	userId, err := h.Services.User.Create(c, body.Email, body.Password)
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

	c.JSON(http.StatusCreated, gin.H{
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
		c.AbortWithStatusJSON(http.StatusNotFound, h.newError("NOT_FOUND"))
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
		c.AbortWithStatusJSON(http.StatusUnauthorized, h.newError("UNAUTHORIZED"))
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

	proceed, err := h.verifyRequest(c, verification.Recovery, body.Email, body.Verification)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	} else if !proceed {
		return
	}

	err = h.Services.User.UpdatePassword(c, user.Id, body.Password)
	if errors.Is(err, user_service.ErrInvalidPassword) {
		c.AbortWithStatusJSON(http.StatusBadRequest, h.newError("INVALID_PASSWORD"))
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
