package handler

import (
	"net/http"
	"strings"

	"github.com/Beksultan15/project_go"
	"github.com/gin-gonic/gin"
)
func (h *Handler) validate(c *gin.Context,input project_go.User) (int){

	if input.Password != input.Confirmpassword {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Passwords do not match"})
		return 0
	}

	if !strings.Contains(input.Email, "@") {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Email doesn't have @"})
		return 0
	}

	return 1
}