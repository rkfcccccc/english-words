package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) GetChallenge(c *gin.Context) {
	userId := c.GetInt("user_id")

	challenge, err := h.Services.Vocabulary.GetChallenge(c, userId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Vocabulary.GetChallenge: %v", err))
		return
	}

	entry, err := h.Services.Dictionary.GetById(c, challenge.WordId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Dictionary.GetById: %v", err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"entry":         entry,
		"learning_step": challenge.LearningStep,
	})
}
