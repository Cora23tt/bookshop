package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *HandlerV1) ReturnBadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": message,
	})
}
