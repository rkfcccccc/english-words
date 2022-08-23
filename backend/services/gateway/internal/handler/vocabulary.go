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

type finishChallengeInput struct {
	Action string `json:"action" binding:"required"`
	WordId string `json:"word_id" binding:"required"`
}

func (h *Handlers) FinishChallenge(c *gin.Context) {
	var body finishChallengeInput

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userId := c.GetInt("user_id")

	var err error
	switch body.Action {
	case "promote":
		err = h.Services.Vocabulary.PromoteWord(c, userId, body.WordId)
	case "resist":
		err = h.Services.Vocabulary.ResistWord(c, userId, body.WordId)
	default:
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("got unknown action: %s", body.Action))
		return
	}

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}

type alreadyLearnedInput struct {
	WordId string `json:"word_id" binding:"required"`
	State  bool   `json:"state" binding:"required"`
}

func (h *Handlers) AlreadyLearned(c *gin.Context) {
	var body alreadyLearnedInput

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userId := c.GetInt("user_id")
	err := h.Services.Vocabulary.SetAlreadyLearned(c, userId, body.WordId, body.State)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}
