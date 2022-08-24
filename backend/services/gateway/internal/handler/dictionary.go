package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) DictionarySearch(c *gin.Context) {
	query := c.Query("query")

	entries, err := h.Services.Dictionary.Search(c, query)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, entries)
}
