package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rkfcccccc/english_words/services/gateway/internal/service/verification"
)

type verificationInput struct {
	RequestId string `json:"request_id"`
	Code      int    `json:"code"`
}

func (h *Handlers) verifyRequest(c *gin.Context, typeId verification.Type, email string, data *verificationInput) (bool, error) {
	if data == nil {
		requestId, err := h.Services.Verification.SendCode(c, email, typeId)
		if err != nil {
			return false, fmt.Errorf("Verification.SendCode: %v", err)
		}

		c.JSON(http.StatusAccepted, gin.H{"request_id": requestId})
		return false, nil
	}

	success, err := h.Services.Verification.Verify(c, data.RequestId, data.Code)
	if err != nil {
		return false, fmt.Errorf("Verification.Verify: %v", err)
	}

	return success, err
}
