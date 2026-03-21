package handlers

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/Artmoond/Minion-Team-TTK-Case/Backend/administration_panel/internal/entity/custom_err"
	"github.com/Artmoond/Minion-Team-TTK-Case/Backend/administration_panel/internal/entity/models"
	"github.com/gin-gonic/gin"
)

func (h *handlers) GetAllUsers(c *gin.Context) {
	var req *models.GetAllUsersRequest

	authHeader := c.GetHeader("Authorization")

	if len(authHeader) == 0 {
		log.Println("err auth header is empty: ", authHeader)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing authorization header"})
		return
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header"})
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	tokenString = strings.TrimSpace(tokenString)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.Token = tokenString

	resp, err := h.service.GetAllUsers(c.Request.Context(), req)
	if err != nil {
		switch {
		case errors.Is(err, custom_err.ErrNotHaveRightRole):
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		case errors.Is(err, custom_err.ErrTokenInvalid):
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		case errors.Is(err, custom_err.ErrEmptyToken):
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		case errors.Is(err, custom_err.ErrGetClaims):
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": resp})
}
